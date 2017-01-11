/* vim: set expandtab tabstop=4 shiftwidth=4 softtabstop=4: */

/**
 * jquery.WebSocket
 *
 * jquery.WebSocket.js - enables WebSocket support with emulated fallback
 *
 * One simple interface $.WebSocket(url, protocol, options); thats it.
 * The same interface as current native WebSocket implementation. The same
 * native events (onopen, onmessage, onerror, onclose) + custom event onsend.
 *
 * But jquery.WebSockets adds some nice features:
 *
 *  [x] Multiplexing - Use a single socket connection and as many logical pipes
 *      within as your browser supports. All these pipes are emulated WebSockets
 *      also with the same API + same events! Use each pipe as WebSocket! But
 *      this requires you to implement the protocol on this level of communication
 *      The data is en- + decoded in a special way to make multiplexing possible
 *
 *  [x] Interface for adding protocol to manipulate data before they are send
 *      and right after they arrive before event onmessage is fired!
 *
 * LICENSE:
 * jquery.WebSocket
 *
 * Copyright (c) 2012, Benjamin Carl - All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 * - All advertising materials mentioning features or use of this software
 *   must display the following acknowledgement: This product includes software
 *   developed by Benjamin Carl and other contributors.
 * - Neither the name Benjamin Carl nor the names of other contributors
 *   may be used to endorse or promote products derived from this
 *   software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 * Please feel free to contact us via e-mail: opensource@clickalicious.de
 *
 * @category   jquery
 * @package    jquery_plugin
 * @subpackage jquery_plugin_WebSocket
 * @author     Benjamin Carl <opensource@clickalicious.de>
 * @copyright  2012 - 2013 Benjamin Carl
 * @license    http://www.opensource.org/licenses/bsd-license.php The BSD License
 * @version    0.0.3
 * @link       http://www.clickalicious.de
 * @see        -
 * @since      File available since Release 1.0.0
 */
(function($){

    // attach to jQuery
    $.extend({

        // export WebSocket = $.WebSocket
        WebSocket: function(url, protocol, options)
        {
            /**
             * The default values, use comma to separate the settings,
             * example:
             *
             * @type {jquery.WebSocket}
             * @private
             */
            var defaults = {
                url: url,
                http: null,
                enableProtocols: false,
                enablePipes: false,
                encoding: 'utf-8',                                                // fallback: encoding for AJAX LP
                method: 'post',                                                   // fallback: method for AJAX LP
                delay: 0,                                                         // number of ms to delay open event
                interval: 3000                                                    // number of ms between poll request
            };

            // overwrite (append) to option defaults
            options = $.extend(
                {},
                defaults,
                options
            );

            // WebSocket Id and readyStates
            const WS_ID      = 'WebSocketPipe';
            const CONNECTING = 0;
            const OPEN       = 1;
            const CLOSING    = 2;
            const CLOSED     = 3;

            // the function table to store references to callbacks
            var _functionTable = {
                onopen:    function()  {},
                onerror:   function(e) {},
                onclose:   function()  {},
                onmessage: function(e) {},
                send:      function(d) { _ws._send(d); }
            };

            /***********************************************************************************************************
             *
             * PRIVATE MEMBERS
             *
             **********************************************************************************************************/

            /**
             * private: _token
             *
             * Returns a random token on request
             *
             * This method is intend to return a random token on
             * request e.g. used as pipe-Id
             *
             * @returns {String} Random token
             */
            function _token()
            {
                return Math.random().toString(36).substr(2);
            };

            /**
             * private: _urlWsToHttp
             *
             * Converts a ws:// or wss:// style link to a http:// or https:// link
             *
             * This method is intend to convert ws-links to http-links
             *
             * @returns {String} http-link
             */
            function _urlWsToHttp(url)
            {
                var protocol = (url.attr('protocol') === 'wss') ? 'https://' : 'http://';
                var host     = url.attr('host');
                var port     = (
                                   (url.attr('protocol') == 'wss' && url.attr('port') != 443) ||
                                   (url.attr('protocol') == 'ws' && url.attr('port') != 80) ?
                                   ':' + url.attr('port') : ''
                               );
                var path     = ((url.attr('path') != '/') ? url.attr('path') : '');

                // return new combined url
                return protocol + host + port + path;
            };

            /**
             * private: _dispatchProtocol
             *
             * Dispatch event to registered protocol handler
             * (events: onmessage, onsend[only on emulated WebSocket!])
             *
             * @param e The event object
             *
             * @returns The processed object
             */
            function _dispatchProtocol(e, direction, ws)
            {
                // give object event from protocol to protocol
                for (var protocol in ws.protocols) {
                    e = ws.protocols[protocol].callback(e, direction);
                }

                // return dirty object
                return e;
            };

            /**
             * private: _dispatchPipe
             *
             *  Dispatch event to pipe
             * (events: onmessage, onsend)
             */
            function _dispatchPipe(id, e)
            {
                for (var pipe in ws.pipes) {
                    if (pipe == id) {
                        var p = ws.pipes[pipe];
                        if (p.onmessage !== undefined &&
                                typeof(p.onmessage) === 'function'
                            ) {
                               p.onmessage(e);
                        }
                    }
                }
            };

            /**
             * private: _proxy
             *
             * This is the proxy between an intercepted call and calls
             * the user defined callback
             *
             * @param {String} trigger The trigger (event-name) to dispatch
             * @param {Object} e       The event object to dispatch
             *
             * @returns The result of dispatch (depends on operation!)
             * @private
             */
            function _proxy(trigger, e)
            {
                return _functionTable[trigger](e);
            };

            /**
             * private: _injectHook
             *
             * Injects hooks into a given WebSocket object
             *
             * This method injects the hooks for the given name of an event (event)
             * into a given WebSocket object (ws).
             *
             * @param {String} event The name of the event to hook
             * @param {Object} ws    An instance of a WebSocket to patch (inject to)
             *
             * @returns void
             * @private
             */
            function _injectHook(event, ws)
            {
                //
                if (event == 'send') {
                    ws._send = ws.send;
                }

                // install our proxy method for intercepting events
                ws[event] = function(e) {
                    var uid;

                    if (event == 'onmessage') {
                        // apply registered protocols first
                        e = _dispatchProtocol(e, 'i', ws);

                        // where to get info -> multiplexed?
                        // if multiplexed look into package for uid
                        if (ws.multiplexed) {
                            // extract uid
                            uid = /"uid"\:\s"([a-z0-9A-z]*)"/.exec(e.data)[1];
                        }

                        if (uid !== undefined) {
                            $(ws.pipes[uid]).trigger(e);
                        } else {
                            // dispatch to all pipes
                            for (pipe in ws.pipes) {
                                $(ws.pipes[pipe]).trigger(
                                    e
                                );
                            }
                        }
                    }

                    (event == 'send') ? e = _dispatchProtocol(e, 'o', ws) : null;

                    _proxy(event, e);
                };

                // override setter with custom hook to fetch user defined callbacks
                window.WebSocket.prototype.__defineSetter__(
                   event,
                    function(v) {
                        _functionTable[event] = v;
                    }
                );

                // override getter with custom hook to fetch user defined callbacks
                window.WebSocket.prototype.__defineGetter__(
                    event,
                    function()  {
                        return _functionTable[event];
                    }
                );
            };

            /**
             * Returns a fresh pipe-object, which is in fact an emulated WebSocket
             * which is extended here within on the fly with our pipe logic.
             *
             * @param {String} url The url connect to (resource or path ...)
             * @param {String} id The optional Id of the pipe (unique-identifier)
             *
             * @returns {Object} A fresh pipe (emulated WebSocket)
             * @private
             */
            function _protocol(name, callback)
            {
                return {
                    name: name,
                    callback: callback
                };
            };

            /**
             * Returns a fresh pipe-object, which is in fact an emulated WebSocket
             * which is extended here within on the fly with our pipe logic.
             *
             * @param {String} url The url connect to (resource or path ...)
             * @param {String} id The optional Id of the pipe (unique-identifier)
             *
             * @returns {Object} A fresh pipe (emulated WebSocket)
             * @private
             */
            function _pipe(url)
            {
                // create, merge and return new pipe instance
                return $.extend(
                    _getWebSocketSkeleton(url, false),     // merge an websocket structure  _getWebSocketSkeleton(url)
                    {                                      // with our additions            {}
                        packet: function(uid) {
                            return {
                                type: WS_ID,
                                action: 'message',
                                data: null,
                                uid: uid
                            };
                        },

                        startup: function() {
                            var p = this.packet(this.uid, this.url);
                            p.action = 'startup';
                            p.data   = this.url;

                            _ws.send(
                                JSON.stringify(p)
                            );

                        },

                        shutdown: function() {
                            // send shutdown pipe packet
                            var p = this.packet(this.uid, this.url);
                            p.action = 'shutdown';

                            // send as text string
                            _ws.send(
                                JSON.stringify(p)
                            );
                        },

                        send: function(data) {
                            // get packet
                            var p = this.packet(this.uid, this.url);
                            p.data = data;

                            // send as text string
                            _ws.send(
                                JSON.stringify(p)
                            );

                            // trigger event send
                            $(this).triggerHandler('send');
                        },

                        start: function() {
                            // send initial packet to server for handshaking
                            this.startup();

                            // trigger event open
                            $(this).triggerHandler('open');

                            return this;
                        },

                        close: function() {
                            // send packet close
                            this.shutdown();

                            // trigger native event WebSocket:onclose (native)
                            $(this).triggerHandler('close');
                        }
                    }
                );
            }

            /**
             * Returns skeleton of a WebSocket object
             *
             * @param {String} url The url connect to (resource or path ...)
             * @param {String} id  The optional Id of the pipe (unique-identifier)
             *
             * @returns {Object} A fresh pipe (emulated WebSocket)
             * @private
             */
            function _getWebSocketSkeleton(url, isNative)
            {
                return {
                    type: 'WebSocket',                  // CUSTOM type     String type of object
                    uid: _token(),                      // CUSTOM uid      an unique Id used to identify the ws
                    //native: isNative,                   // CUSTOM native   status as boolean
                    readyState: CONNECTING,             // NATIVE default  status is 0 = CONNECTING
                    bufferedAmount: 0,                  // NATIVE integer  currently buffered data in bytes
                    url: url,                           // NATIVE url      The url to connect to

                    send: function(data) {},            // NATIVE send()   sending "data" to server
                    start: function() {},               // CUSTOM start()  used as custom trigger for opening
                    close: function() {},               // NATIVE close()  closes the connection

                    onopen: function() {},              // NATIVE EVENT
                    onerror: function(e) {},            // NATIVE EVENT
                    onclose: function() {},             // NATIVE EVENT
                    onmessage: function(e) {},          // NATIVE EVENT

                    protocols: {},                      // CUSTOM          container for protocols in user defined order
                    pipes: {},                          // CUSTOM          container for all registered pipes
                    multiplexed: false,                 // CUSTOM          multiplexed status of this object

                    /**
                     * Export: registerPipe()
                     *
                     * Creates a new pipe and return pipe-reference.
                     *
                     * This method creates a "pipe" which is in fact an emulated WebSocket
                     * (the same emulation as used in older browsers). These "pipes" are
                     * used as a logical connection within a physical connection to the Server
                     *
                     * Pipe 1 -> O==\     (WebSocket)    /==O -> Pipe 1
                     * Pipe 2 -> O==========================O -> Pipe 2
                     * Pipe 3 -> O==/                    \==O -> Pipe 3
                     *
                     * So we can use those Pipes in the same way like the original WebSocket object.
                     * You can use the same events what enables you and your app to use less
                     * connections between server and client and use different endpoints for your
                     * services too.
                     *
                     * @param {String} url      The url to connect to
                     * @param {String} protocol The protocol to use
                     * @param {Object} options  Custom options to pass through
                     *
                     * @returns {Object} the instance created
                     */
                    registerPipe: function(url, protocol, options) {
                        var p = new _pipe(url);
                        this.multiplexed = true;

                        // we iterate the functionTable and use the events for injecting our hooks
                        for (event in _functionTable) {
                            // inject all hooks except "send"
                            if (event != 'send') {
                                _injectHook(event, p);
                            }
                        }

                        // try:
                        // return ws.pipes[p.id] = p = new _pipe(url, id);
                        var r = this.pipes[p.uid] = $.extend(p, options).start();
                        return r;
                    },
                    unregisterPipe: function(id) {
                        this.pipes[id] = null;
                        (this.pipes.length === 0) ? this.multiplexed = false : void(0);
                        // timer?
                    },
                    registerProtocol: function(name, callback) {
                        var p = new _protocol(name, callback);
                        return this.protocols[name] = p;
                    },
                    unregisterProtocol: function(name) {
                        this.protocols[name] = null;
                    },
                    extension: null,
                    protocol: null,
                    reason: null,
                    binaryType: null
                };
            }

            /**
             * Creates and return a new jQuery error event with passed var added as .data
             *
             * @param {Mixed} data The data to add to event
             *
             * @returns {jQuery.Event} An jQuery error event object
             * @private
             */
            function _ErrorEvent(data)
            {
                // create MessageEvent event and add received data
                var e = jQuery.Event('error');
                e.data = data;
                return e;
            }

            /**
             * Creates and return a new jQuery message event with passed var added as .data
             *
             * @param {Mixed} data The data to add to event
             *
             * @returns {jQuery.Event} An jQuery message event object
             * @private
             */
            function _MessageEvent(data)
            {
                // create MessageEvent event and add received data
                var e = jQuery.Event('message');
                e.data = data;
                return e;
            }

            /**
             * Returns an emulated WebSocket which is build upon jQueries AJAX functionality.
             *
             * @param {Mixed} data The data to add to event
             *
             * we create and return an emulated WebSocket Object. We you can use this object in a very
             * similar way to the native WebSocket Implementation. The connection was graceful degraded to
             * AJAX long polling ...
             *
             * @returns ???
             * @private
             */
            function _WebSocket(url)
            {
                var _interval, _handler, _emulation = {

                    /**
                     * export: send
                     *
                     * Sends data via options.method to server (http/xhr request)
                     *
                     * This method is intend to ...
                     *
                     * @param mixed data The data to send
                     *
                     * @returns boolean TRUE on success, otherwise FALSE
                     */
                    send: function(data) {
                        // default result is true = success
                        var success = true;

                        // send data via jQuery ajax()
                        $.ajax({
                            async: false,
                            type: options.method,
                            url: url + (
                                (options.method == 'GET' && options.arguments) ?
                                 '?' + $.param(options.arguments) :
                                 ''
                            ),
                            data: (
                                (options.method == 'POST' && options.arguments) ?
                                 $.param($.extend(options.arguments, {data: data})) :
                                 null
                            ),
                            dataType: 'text',
                            contentType: "application/x-www-form-urlencoded; charset=" + options.encoding,
                            success: function(data) {
                                // trigger native event MessageEvent (emulated)
                                $(_emulation).trigger(
                                    new _MessageEvent(_dispatchProtocol(data, 'i', _emulation))
                                );
                            },
                            error: function(xhr, data, errorThrown) {
                                // in case of error no success
                                success = false;

                                // trigger native event ErrorEvent (emulated)
                                $(_emulation).trigger(
                                    _ErrorEvent(data)
                                );
                            }
                        });

                        // return result of operation
                        return success;
                    },

                    /**
                     * export: close
                     *
                     * Closes an existing and open connection
                     *
                     * This method is intend to ...
                     *
                     * @returns void
                     */
                    close: function() {
                        // kill timer!
                        clearTimeout(_handler);
                        clearInterval(_interval);

                        // update readyState
                        this.readyState = CLOSED;

                        // trigger native event WebSocket:onclose (native)
                        $(_emulation).triggerHandler('close');
                    }
                };

                /**
                 * private: _poll
                 *
                 * Polls server for new data and returns the result
                 *
                 * This method is intend to ...
                 *
                 * @returns void
                 * @private
                 */
                function _poll() {
                    $.ajax({
                        type: options.method,
                        url: url + (
                            (options.method == 'GET' && options.arguments) ?
                             '?' + $.param(options.arguments) :
                             ''
                        ),
                        dataType: 'text',
                        data: (
                            (options.method == 'POST' && options.arguments) ?
                             $.param(options.arguments) :
                             null
                        ),
                        success: function(data) {
                            // trigger our emulated MessageEvent
                            $(_emulation).trigger(
                                new _MessageEvent(data)
                            );
                        },
                        error: function(xhr, data, errorThrown) {
                            // in case of error no success
                            success = false;

                            // trigger native event ErrorEvent (emulated)
                            $(_emulation).trigger(
                                _ErrorEvent(data)
                            );
                        }
                    });

                    // trigger custom event WebSocket:onsend (custom)
                    $(_emulation).triggerHandler('send');
                };

                // run our emulation
                _handler = setTimeout(
                    function() {
                        _emulation.readyState = OPEN;
                        _poll();
                        _interval = setInterval(_poll, options.interval);

                        // trigger event WebSocket:onopen (emulated)
                        $(_emulation).triggerHandler('open');
                    },
                    options.delay
                );

                // return emulated socket implementation
                return _emulation;
            };

            /**
             * private: _extend
             *
             * Copies all properties from source to destination object as
             * long as they not exist in destination.
             *
             * This method is intend to ...
             *
             * @param {Object} source      The object to copy from
             * @param {Object} destination The object to copy to
             *
             * @returns {Object} The object destination with new properties set
             * @private
             */
            function _extend(source, destination) {
              for (property in source) {
                if (!destination[property]) {
                  destination[property] = source[property];
                }
              }

              return destination;
            }

            /**
             * private: _getWebSocket
             *
             * Returns a WebSocket-Object (native or emulated)
             *
             * This method is intend to return a WebSocket-Object no matter
             * if the browser supports native WebSockets or not. For older
             * Browsers the WebSocket-Object is emulated with Ajax (Fake)
             * Push: Long Polling.
             *
             * @param {String} url The resource-locator (e.g. ws://127.0.0.1:80/ ||
             *                        ws://127.0.0.1:80/this/is/fallback/)
             *
             * @returns {Object} WebSocket
             * @private
             */
            function _getWebSocket(url)
            {
              var ws, isNative = true;

                // websocket support built in?
                if (window.WebSocket) {
                    // in firefox we use MozWebSocket (but it seems that from now (FF 17) MozWebsocket is removed)
                    if (typeof(MozWebSocket) == 'function') {
                      ws = new MozWebSocket(url);
                    } else {
                        ws = new WebSocket(url);
                    }

                } else {
                    // inject emulated WebSocket implementation into DOM
                    window.WebSocket = _WebSocket;

                    url      = options.http;
                    isNative = false;
                    ws       = new WebSocket(url);
                }

                // extend it with our additions and return
                return _extend(
                    _getWebSocketSkeleton(url, isNative),
                    ws
                );
            };

            /***********************************************************************************************************
             *
             *  MAIN()
             *
             **********************************************************************************************************/

            // get WebSocket object (either native or emulated)
            var _ws = _getWebSocket(options.url);

            // we iterate the functionTable and use the events for injecting our hooks
            for (event in _functionTable) {
                _injectHook(event, _ws);
            }

            // return WebSocket
            return _ws;
        }
    });
})(jQuery);
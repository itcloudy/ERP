/*! Bootstrap integration for DataTables' Editor
 * ©2015 SpryMedia Ltd - datatables.net/license
 */

(function( factory ){
	if ( typeof define === 'function' && define.amd ) {
		// AMD
		define( ['jquery', 'datatables.net-bs', 'datatables.net-editor'], function ( $ ) {
			return factory( $, window, document );
		} );
	}
	else if ( typeof exports === 'object' ) {
		// CommonJS
		module.exports = function (root, $) {
			if ( ! root ) {
				root = window;
			}

			if ( ! $ || ! $.fn.dataTable ) {
				$ = require('datatables.net-bs')(root, $).$;
			}

			if ( ! $.fn.dataTable.Editor ) {
				require('datatables.net-editor')(root, $);
			}

			return factory( $, root, root.document );
		};
	}
	else {
		// Browser
		factory( jQuery, window, document );
	}
}(function( $, window, document, undefined ) {
'use strict';
var DataTable = $.fn.dataTable;


/*
 * Set the default display controller to be our bootstrap control 
 */
DataTable.Editor.defaults.display = "bootstrap";


/*
 * Alter the buttons that Editor adds to TableTools so they are suitable for bootstrap
 */
var i18nDefaults = DataTable.Editor.defaults.i18n;
i18nDefaults.create.title = "<h3>"+i18nDefaults.create.title+"</h3>";
i18nDefaults.edit.title = "<h3>"+i18nDefaults.edit.title+"</h3>";
i18nDefaults.remove.title = "<h3>"+i18nDefaults.remove.title+"</h3>";

var tt = DataTable.TableTools;
if ( tt ) {
	tt.BUTTONS.editor_create.formButtons[0].className = "btn btn-primary";
	tt.BUTTONS.editor_edit.formButtons[0].className = "btn btn-primary";
	tt.BUTTONS.editor_remove.formButtons[0].className = "btn btn-danger";
}


/*
 * Change the default classes from Editor to be classes for Bootstrap
 */
$.extend( true, $.fn.dataTable.Editor.classes, {
	"header": {
		"wrapper": "DTE_Header modal-header"
	},
	"body": {
		"wrapper": "DTE_Body modal-body"
	},
	"footer": {
		"wrapper": "DTE_Footer modal-footer"
	},
	"form": {
		"tag": "form-horizontal",
		"button": "btn btn-default"
	},
	"field": {
		"wrapper": "DTE_Field",
		"label":   "col-lg-4 control-label",
		"input":   "col-lg-8 controls",
		"error":   "error has-error",
		"msg-labelInfo": "help-block",
		"msg-info":      "help-block",
		"msg-message":   "help-block",
		"msg-error":     "help-block",
		"multiValue":    "well well-sm multi-value",
		"multiInfo":     "small",
		"multiRestore":  "well well-sm multi-restore"
	}
} );


/*
 * Bootstrap display controller - this is effectively a proxy to the Bootstrap
 * modal control.
 */

var self;

DataTable.Editor.display.bootstrap = $.extend( true, {}, DataTable.Editor.models.displayController, {
	/*
	 * API methods
	 */
	"init": function ( dte ) {
		// init can be called multiple times (one for each Editor instance), but
		// we only support a single construct here (shared between all Editor
		// instances)
		if ( ! self._dom.content ) {
			self._dom.content = $(
				'<div class="modal fade">'+
					'<div class="modal-dialog">'+
						'<div class="modal-content"/>'+
					'</div>'+
				'</div>'
			);

			self._dom.close = $('<button class="close">&times;</div>');

			self._dom.close.click( function () {
				self._dte.close('icon');
			} );

			$(document).on('click', 'div.modal', function (e) {
				if ( $(e.target).hasClass('modal') && self._shown ) {
					self._dte.background();
				}
			} );

			dte.on( 'open.dtebs', function ( e, type ) {
				if ( type === 'inline' || type === 'bubble' ) {
					$('div.DTE input[type=text], div.DTE select, div.DTE textarea').addClass( 'form-control' );
				}
			} );
		}

		return self;
	},

	"open": function ( dte, append, callback ) {
		if ( self._shown ) {
			if ( callback ) {
				callback();
			}
			return;
		}

		self._dte = dte;
		self._shown = true;

		var content = self._dom.content.find('div.modal-content');
		content.children().detach();
		content.append( append );

		$('div.modal-header', append).prepend( self._dom.close );

		$(self._dom.content)
			.one('shown.bs.modal', function () {
				// Can only give elements focus when shown
				if ( self._dte.s.setFocus ) {
					self._dte.s.setFocus.focus();
				}

				if ( callback ) {
					callback();
				}
			})
			.one('hidden', function () {
				self._shown = false;
			})
			.appendTo( 'body' )
			.modal( {
				backdrop: "static",
				keyboard: false
			} );

		$('input:not([type=checkbox]):not([type=radio]), select, textarea', self._dom.content)
			.addClass( 'form-control' );
	},

	"close": function ( dte, callback ) {
		if ( !self._shown ) {
			if ( callback ) {
				callback();
			}
			return;
		}

		$(self._dom.content)
			.one( 'hidden.bs.modal', function () {
				$(this).detach();
			} )
			.modal('hide');

		self._dte = dte;
		self._shown = false;

		if ( callback ) {
			callback();
		}
	},

	node: function ( dte ) {
		return self._dom.content[0];
	},


	/*
	 * Private properties
	 */
	 "_shown": false,
	"_dte": null,
	"_dom": {}
} );

self = DataTable.Editor.display.bootstrap;


return DataTable.Editor;
}));

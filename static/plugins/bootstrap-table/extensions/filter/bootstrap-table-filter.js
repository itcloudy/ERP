!function($) {

    'use strict';

    // TOOLS DEFINITION
    // ======================
    var rowLabel = function(el) {
        var ret = el;
        if (typeof el === 'object') {
            ret = el.label;
            if (typeof el.i18n === 'object') {
                $.each(el.i18n, function(key, val) { ret = ret.replace('{%' + key + '}', val) });
            }
        }
        return ret;
    };
    var rowId = function(id, el) {
        return typeof el === 'object' ? el.id : id;
    };
    var getOptionData = function($option) {
        var val = false;
        var name;
        var data = {}, cnt = 0;
        var $chck = $option.find('.filter-enabled');
        $(':input', $option).each(function() {
            var $this = $(this);
            if ($this.is($chck)) {
                return;
            }
            name = $this.attr('data-name');
            if (name) {
                data[name] = $this.val();
            }
            val = $this.val();
            cnt++;
        });
        return $.isEmptyObject(data) ? val : data;
    };


    // FILTER CLASS DEFINITION
    // ======================

    var BootstrapTableFilter = function(el, options) {
        this.options = options;
        this.$el = $(el);
        this.$el_ = this.$el.clone();
        this.timeoutId_ = 0;
        this.filters = {};

        this.init();
    };

    BootstrapTableFilter.DEFAULTS = {
        filters: [],
        connectTo: false,

        filterIcon: '<span class="glyphicon glyphicon-filter"></span>',
        refreshIcon: '<span class="glyphicon glyphicon-ok"></span>',
        clearAllIcon: '<span class="glyphicon glyphicon-remove"></span>',

        formatRemoveFiltersMessage: function() {
            return 'Remove all filters';
        },
        formatSearchMessage: function() {
            return 'Search';
        },

        onAll: function(name, args) {
            return false;
        },
        onFilterChanged: function(data) {
            return false;
        },
        onResetView: function() {
            return false;
        },
        onAddFilter: function(filter) {
            return false;
        },
        onRemoveFilter: function(field) {
            return false;
        },
        onEnableFilter: function(field) {
            return false;
        },
        onDisableFilter: function(field) {
            return false;
        },
        onSelectFilterOption: function(field, option, data) {
            return false;
        },
        onUnselectFilterOption: function(field, option) {
            return false;
        },
        onDataChanged: function(data) {
            return false;
        },
        onSubmit: function(data) {
            return false;
        },
    };

    BootstrapTableFilter.EVENTS = {
        'all.bs.table.filter': 'onAll',
        'reset.bs.table.filter': 'onResetView',
        'add-filter.bs.table.filter': 'onAddFilter',
        'remove-filter.bs.table.filter': 'onRemoveFilter',
        'enable-filter.bs.table.filter': 'onEnableFilter',
        'disable-filter.bs.table.filter': 'onDisableFilter',
        'select-filter-option.bs.table.filter': 'onSelectFilterOption',
        'unselect-filter-option.bs.table.filter': 'onUnselectFilterOption',
        'data-changed.bs.table.filter': 'onDataChanged',
        'submit.bs.table.filter': 'onSubmit'
    };

    BootstrapTableFilter.FILTER_SOURCES = {
        range: {
            search: false,
            rows: [
                {id: 'lte', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Less than'}},
                {id: 'gte', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'More than'}},
                {id: 'eq', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Equals'}}
            ],
            check: function(filterData, value) {
                if (typeof filterData.lte !== 'undefined' && parseInt(value) > parseInt(filterData.lte)) {
                    return false;
                }
                if (typeof filterData.gte !== 'undefined' && parseInt(value) < parseInt(filterData.gte)) {
                    return false;
                }
                if (typeof filterData.eq !== 'undefined' && parseInt(value) != parseInt(filterData.eq)) {
                    return false;
                }
                return true;
            }
        },
        search: {
            search: false,
            rows: [
                {id: 'eq', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Equals'}},
                {id: 'neq', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Not equals'}},
                {id: 'cnt', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Contains'}},
                {id: 'ncnt', label: '{%msg} <input class="form-control" type="text">', i18n: {msg: 'Doesn\'t contain'}},
                {id: 'ept', label: '{%msg}', i18n: {msg: 'Is empty'}},
                {id: 'nept', label: '{%msg}', i18n: {msg: 'Is not empty'}}
            ],
            check: function(filterData, value) {
                if (typeof filterData.eq !== 'undefined' && value != filterData.eq) {
                    return false;
                }
                if (typeof filterData.neq !== 'undefined' && value == filterData.neq) {
                    return false;
                }
                if (typeof filterData.cnt !== 'undefined' && value.indexOf(filterData.cnt) < 0) {
                    return false;
                }
                if (typeof filterData.ncnt !== 'undefined' && value.indexOf(filterData.ncnt) >= 0) {
                    return false;
                }
                if (typeof filterData._values !== 'undefined' && filterData._values.indexOf('ept') >= 0 && value.trim()) {
                    return false;
                }
                if (typeof filterData._values !== 'undefined' && filterData._values.indexOf('nept') >= 0 && !value.trim()) {
                    return false;
                }
                return true;
            }
        },
        ajaxSelect: {
            search: true,
            rows: [],
            rowsCallback: function(filter, searchPhrase) {
                var that = this;
                clearTimeout(this.timeoutId_);
                this.timeoutId_ = setTimeout(function() {
                    $.ajax(filter.source, {dataType: 'json', data: {q: searchPhrase}})
                    .done(function(data) {
                        that.clearFilterOptions(filter.field);
                        that.fillFilterOptions(filter.field, data);
                    });
                }, 300);
            }
        },
        select: {
            search: true,
            rows: [],
            rowsCallback: function(filter, searchPhrase) {
                var vals = filter.values;
                var label;
                if (searchPhrase.length) {
                    vals = vals.filter(function(el) {
                        return rowLabel(el).indexOf(searchPhrase) > -1
                    });
                }
                this.clearFilterOptions(filter.field);
                this.fillFilterOptions(filter.field, vals.slice(0, 20));
            }
        }
    };

    BootstrapTableFilter.EXTERNALS = [];

    BootstrapTableFilter.prototype.init = function() {
        this.initContainer();
        this.initMainButton();
        this.initFilters();
        this.initRefreshButton();
        this.initFilterSelector();
        this.initExternals();
    };

    BootstrapTableFilter.prototype.initContainer = function() {
        var that = this;
        this.$toolbar = $([
            '<div class="btn-toolbar">',
                '<div class="btn-group btn-group-filter-main">',
                    '<button type="button" class="btn btn-default dropdown-toggle btn-filter" data-toggle="dropdown">',
                        this.options.filterIcon,
                    '</button>',
                    '<ul class="dropdown-menu" role="menu">',
                    '</ul>',
                '</div>',
                '<div class="btn-group btn-group-filters">',
                '</div>',
                '<div class="btn-group btn-group-filter-refresh">',
                    '<button type="button" class="btn btn-default btn-primary btn-refresh" data-toggle="dropdown">',
                        this.options.refreshIcon,
                    '</button>',
                '</div>',
            '</div>'
        ].join(''));
        this.$toolbar.appendTo(this.$el);
        this.$filters = this.$toolbar.find('.btn-group-filters');

        this.$toolbar.delegate('.btn-group-filters li', 'click', function (e) {
            e.stopImmediatePropagation();
        });

        this.$toolbar.delegate('.btn-group-filters li .filter-enabled', 'click', function(e) {
            var $chck = $(this);
            var field = $chck.closest('[data-filter-field]').attr('data-filter-field');
            var $option = $chck.closest('[data-val]');
            var option = $option.attr('data-val');
            if ($chck.prop('checked')) {
                var data = getOptionData($option);
                that.selectFilterOption(field, option, data);
            }
            else {
                that.unselectFilterOption(field, option);
            }
            e.stopImmediatePropagation();
        });
        this.$toolbar.delegate('.btn-group-filters li :input:not(.filter-enabled)', 'click change', function(e) {
            var $inp = $(this);
            var field = $inp.closest('[data-filter-field]').attr('data-filter-field');
            var $option = $inp.closest('[data-val]');
            var option = $option.attr('data-val');
            var $chck = $option.find('.filter-enabled');
            if ($inp.val()) {
                var data = getOptionData($option);
                that.selectFilterOption(field, option, data);
                $chck.prop('checked', true);
            }
            else {
                that.unselectFilterOption(field, option);
                $chck.prop('checked', false);
            }
            e.stopImmediatePropagation();
        });
        this.$toolbar.delegate('.search-values', 'keyup', function(e) {
            var $this = $(this);
            var phrase = $this.val();
            var field = $this.closest('[data-filter-field]').attr('data-filter-field');
            var filter = that.getFilter(field);
            var fType = that.getFilterType(filter);
            if (fType.rowsCallback) {
                fType.rowsCallback.call(that, filter, phrase);
            }
        });
    };

    BootstrapTableFilter.prototype.initMainButton = function() {
        this.$button = this.$toolbar.find('.btn-filter');
        this.$buttonList = this.$button.parent().find('.dropdown-menu');
        this.$button.dropdown();
    };

    BootstrapTableFilter.prototype.initRefreshButton = function() {
        var that = this;
        this.$refreshButton = this.$toolbar.find('.btn-refresh');
        this.$refreshButton.click(function(e) {
            that.trigger('submit', that.getData());
            that.toggleRefreshButton(false);
        });
        this.toggleRefreshButton(false);
    };

    BootstrapTableFilter.prototype.initFilters = function() {
        var that = this;
        this.$buttonList.append('<li class="remove-filters"><a href="javascript:void(0)">' + this.options.clearAllIcon + ' ' + this.options.formatRemoveFiltersMessage() + '</a></li>');
        this.$buttonList.append('<li class="divider"></li>');
        $.each(this.options.filters, function(i, filter) {
            that.addFilter(filter);
        });
        this.$toolbar.delegate('.remove-filters *', 'click', function() {
            that.disableFilters();
        });
    };

    BootstrapTableFilter.prototype.initFilterSelector = function() {
        var that = this;
        var applyFilter = function($chck) {
            var filterField = $chck.closest('[data-filter-field]').attr('data-filter-field');
            if ($chck.prop('checked')) {
                that.enableFilter(filterField);
            }
            else {
                that.disableFilter(filterField);
            }
        };
        this.$buttonList.delegate('li :input[type=checkbox]', 'click', function(e) {
            applyFilter($(this));
            e.stopImmediatePropagation();
        });
        this.$buttonList.delegate('li, li a', 'click', function(e) {
            var $chck = $(':input[type=checkbox]', this);
            if ($chck.length) {
                $chck.prop('checked', !$chck.is(':checked'));
                applyFilter($chck);
                e.stopImmediatePropagation();
            }
            var $inp = $(':input[type=text]', this);
            if ($inp.length) {
                $inp.focus();
            }
        });
    };

    BootstrapTableFilter.prototype.initExternals = function() {
        var that = this;
        $.each(BootstrapTableFilter.EXTERNALS, function(i, ext) {
            ext.call(that);
        });
    }

    BootstrapTableFilter.prototype.getFilter = function(field) {
        if (typeof this.filters[field] === 'undefined') {
            throw 'Invalid filter ' + field;
        }
        return this.filters[field];
    };
    BootstrapTableFilter.prototype.getFilterType = function(field, type) {
        if (field) {
            var filter = typeof field === 'object' ? field : this.getFilter(field);
            type = filter.type;
        }
        if (typeof BootstrapTableFilter.FILTER_SOURCES[type] === 'undefined') {
            throw 'Invalid filter type ' + type;
        }
        var ret = BootstrapTableFilter.FILTER_SOURCES[type];
        if (typeof ret.extend !== 'undefined') {
            ret = $.extend({}, ret, this.getFilterType(null, ret.extend));
        }
        return ret;
    };
    BootstrapTableFilter.prototype.checkFilterTypeValue = function(filterType, filterData, value) {
        if (typeof filterType.check === 'function') {
            return filterType.check(filterData, value);
        }
        else {
            if (typeof filterData._values !== 'undefined') {
                return $.inArray("" + value, filterData._values) >= 0;
            }
        }
        return true;
    };

    BootstrapTableFilter.prototype.clearFilterOptions = function(field) {
        var filter = this.getFilter(field);
        filter.$dropdownList.find('li:not(.static)').remove();
        this.toggleRefreshButton(true);
    };

    BootstrapTableFilter.prototype.fillFilterOptions = function(field, data, cls) {
        var that = this;
        var filter = this.getFilter(field);
        cls = cls || '';
        var option, checked;
        $.each(data, function(i, row) {
            option = rowId(i, row);
            checked = that.isSelected(field, option);
            filter.$dropdownList.append($('<li data-val="' + option + '" class="' + cls + '"><a href="javascript:void(0)"><input type="checkbox" class="filter-enabled"' + (checked ? ' checked' : '') + '> ' + rowLabel(row) + '</a></li>'));
        });
    };

    BootstrapTableFilter.prototype.trigger = function(name) {
        var args = Array.prototype.slice.call(arguments, 1);

        name += '.bs.table.filter';
        if (typeof BootstrapTableFilter.EVENTS[name] === 'undefined') {
            throw 'Unknown event ' + name;
        }
        this.options[BootstrapTableFilter.EVENTS[name]].apply(this.options, args);
        this.$el.trigger($.Event(name), args);

        this.options.onAll(name, args);
        this.$el.trigger($.Event('all.bs.table.filter'), [name, args]);
    };

    // PUBLIC FUNCTION DEFINITION
    // =======================

    BootstrapTableFilter.prototype.resetView = function() {
        this.$el.html();
        this.init();
        this.trigger('reset');
    };

    BootstrapTableFilter.prototype.addFilter = function(filter) {
        this.filters[filter.field] = filter;
        this.$buttonList.append('<li data-filter-field="' + filter.field + '"><a href="javascript:void(0)"><input type="checkbox"> ' + filter.label + '</a></li>');

        this.trigger('add-filter', filter);
        if (typeof filter.enabled !== 'undefined' && filter.enabled) {
            this.enableFilter(filter.field);
        }
    };

    BootstrapTableFilter.prototype.removeFilter = function(field) {
        this.disableFilter(field);
        this.$buttonList.find('[data-filter-field=' + field + ']').remove();
        this.trigger('remove-filter', field);
    };

    BootstrapTableFilter.prototype.enableFilter = function(field) {
        var filter = this.getFilter(field);
        var $filterDropdown = $([
            '<div class="btn-group" data-filter-field="' + field + '">',
                '<button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">',
                    filter.label,
                    ' <span class="caret"></span>',
                '</button>',
                '<ul class="dropdown-menu" role="menu">',
                '</ul>',
            '</div>'
        ].join(''));
        $filterDropdown.appendTo(this.$filters);
        filter.$dropdown = $filterDropdown;
        filter.$dropdownList = $filterDropdown.find('.dropdown-menu');
        filter.enabled = true;

        this.$buttonList.find('[data-filter-field=' + field + '] input[type=checkbox]').prop('checked', true);

        var fType = this.getFilterType(filter);
        if (fType.search) {
            filter.$dropdownList.append($('<li class="static"><span><input type="text" class="form-control search-values" placeholder="' + this.options.formatSearchMessage() + '"></span></li>'));
            filter.$dropdownList.append($('<li class="static divider"></li>'));
        }
        if (fType.rows) {
            this.fillFilterOptions(field, fType.rows, 'static');
        }
        if (fType.rowsCallback) {
            fType.rowsCallback.call(this, filter, '');
        }
        this.toggleRefreshButton(true);
        this.trigger('enable-filter', filter);
    };

    BootstrapTableFilter.prototype.disableFilters = function() {
        var that = this;
        $.each(this.filters, function(i, filter) {
            that.disableFilter(filter.field);
        });
    };

    BootstrapTableFilter.prototype.disableFilter = function(field) {
        var filter = this.getFilter(field);
        this.$buttonList.find('[data-filter-field=' + field + '] input[type=checkbox]').prop('checked', false);
        filter.enabled = false;
        if (filter.$dropdown) {
            filter.$dropdown.remove();
            delete filter.$dropdown;
            this.trigger('disable-filter', filter);
        }
        this.toggleRefreshButton(true);
    };

    BootstrapTableFilter.prototype.selectFilterOption = function(field, option, data) {
        var filter = this.getFilter(field);
        if (typeof filter.selectedOptions === 'undefined')
            filter.selectedOptions = {};
        if (data) {
            filter.selectedOptions[option] = data;
        }
        else {
            if (typeof filter.selectedOptions._values === 'undefined') {
                filter.selectedOptions._values = [];
            }
            filter.selectedOptions._values.push(option);
        }
        this.trigger('select-filter-option', field, option, data);
        this.toggleRefreshButton(true);
    };

    BootstrapTableFilter.prototype.unselectFilterOption = function(field, option) {
        var filter = this.getFilter(field);
        if (typeof filter.selectedOptions !== 'undefined' && typeof filter.selectedOptions[option] !== 'undefined') {
            delete filter.selectedOptions[option];
        }
        if (typeof filter.selectedOptions !== 'undefined' && typeof filter.selectedOptions._values !== 'undefined') {
            filter.selectedOptions._values = filter.selectedOptions._values.filter(function(item) {
                return item != option
            });
            if (filter.selectedOptions._values.length == 0) {
                delete filter.selectedOptions._values;
            }
            if ($.isEmptyObject(filter.selectedOptions)) {
                delete filter.selectedOptions;
            }
        }
        this.trigger('unselect-filter-option', field, option);
        this.toggleRefreshButton(true);
    };

    BootstrapTableFilter.prototype.setupFilter = function(field, options) {
        var that = this;
        this.enableFilter(field);
        $.each(options, function(key, val) {
            if (key === '_values') {
                $.each(val, function(i, v) {
                    that.selectFilterOption(field, v, false);
                    $('div[data-filter-field="' + field + '"] [data-val="' + v + '"] input.filter-enabled').prop('checked', true);
                });
            }
            else {
                that.selectFilterOption(field, key, val);
                $('div[data-filter-field="' + field + '"] [data-val="' + key + '"] input.filter-enabled').prop('checked', true);
                $('div[data-filter-field="' + field + '"] [data-val="' + key + '"] input[type="text"]:not([data-name])').val(val);
            }
        });
    };

    BootstrapTableFilter.prototype.toggleRefreshButton = function(show) {
        this.$refreshButton.toggle(show);
    };

    BootstrapTableFilter.prototype.isSelected = function(field, option, value) {
        var filter = this.getFilter(field);
        if (typeof filter.selectedOptions !== 'undefined') {
            if (typeof filter.selectedOptions[option] !== 'undefined') {
                if (value ? (filter.selectedOptions[option] == value) : filter.selectedOptions[option]) {
                    return true
                }
            }
            if (typeof filter.selectedOptions._values !== 'undefined') {
                if (filter.selectedOptions._values.indexOf(option.toString()) > -1) {
                    return true;
                }
            }
        }
        return false;
    };

    BootstrapTableFilter.prototype.getData = function() {
        var that = this;
        var ret = {};
        $.each(that.filters, function(field, filter) {
            if (filter.enabled) {
                if (typeof filter.selectedOptions !== 'undefined') {
                    ret[field] = filter.selectedOptions;
                }
            }
        });
        return ret;
    };

    // BOOTSTRAP FILTER TABLE PLUGIN DEFINITION
    // =======================

    $.fn.bootstrapTableFilter = function(option, _relatedTarget, _param2) {
        BootstrapTableFilter.externals = this.externals;

        var allowedMethods = [
            'addFilter', 'removeFilter',
            'enableFilter', 'disableFilter', 'disableFilters',
            'selectFilterOption', 'unselectFilterOption',
            'setupFilter',
            'toggleRefreshButton',
            'getData', 'isSelected',
            'resetView'
        ],
        value;

        this.each(function() {
            var $this = $(this),
                data = $this.data('bootstrap.tableFilter'),
                options = $.extend(
                    {}, BootstrapTableFilter.DEFAULTS, $this.data(),
                    typeof option === 'object' && option
                );

            if (typeof option === 'string') {
                if ($.inArray(option, allowedMethods) < 0) {
                    throw "Unknown method: " + option;
                }

                if (!data) {
                    return;
                }

                value = data[option](_relatedTarget, _param2);

                if (option === 'destroy') {
                    $this.removeData('bootstrap.tableFilter');
                }
            }

            if (!data) {
                $this.data('bootstrap.tableFilter', (data = new BootstrapTableFilter(this, options)));
            }
        });

        return typeof value === 'undefined' ? this : value;
    };

    $.fn.bootstrapTableFilter.Constructor = BootstrapTableFilter;
    $.fn.bootstrapTableFilter.defaults = BootstrapTableFilter.DEFAULTS;
    $.fn.bootstrapTableFilter.columnDefaults = BootstrapTableFilter.COLUMN_DEFAULTS;
    $.fn.bootstrapTableFilter.externals = BootstrapTableFilter.EXTERNALS;
    $.fn.bootstrapTableFilter.filterSources = BootstrapTableFilter.FILTER_SOURCES;

    // BOOTSTRAP TABLE FILTER INIT
    // =======================

    $(function() {
        $('[data-toggle="tableFilter"]').bootstrapTableFilter();
    });

}(jQuery);

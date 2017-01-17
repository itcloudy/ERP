$(function() {
    'use strict';
    // 产品款式增加属性
    $("#add-one-product-template-attribute").on('click', function(e) {

        var AttributeLineId = "AttributeLineId-" + e.timeStamp;
        var AttributeLineValueId = "AttributeLineValueId-" + e.timeStamp;
        var appendHTML = "<tr data-tree='productAttributes' class='product-template-attribute-line form-tree-create'>" +
            "<td><select data-name='AttributeLineIds' name='AttributeLineId' id='" + AttributeLineId + "' class='form-tree-create form-control select-product-attribute'></select></td>" +
            "<td><select data-name='AttributeLineValueIds'  name='AttributeLineValueId' data-attributelineid='" + AttributeLineId + "' id='" + AttributeLineValueId + "' multiple='multiple' class='form-tree-create form-control select-product-attribute-value'></select></td>" +
            "<td class='text-center'><a class='form-tree-delete'><i class='fa fa-trash-o'></i></a></td>" +
            "</tr>";
        $(appendHTML).prependTo("#product-template-attribute-body");
        $("#" + AttributeLineId).select2({
            width: "off",
            ajax: {
                url: '/product/attribute/?action=search',
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        name: params.term || "", // search term
                        offset: params.page || 0,
                        limit: 5,
                    };
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    // 过滤掉已经添加的属性
                    var existAttr = $("#product-template-attribute-body .select-product-attribute");
                    var exclude = [];
                    for (var i = 0, len = existAttr.length; i < len; i++) {
                        var val = $(existAttr[i]).val();
                        if (val != null) {
                            exclude.push(parseInt(val));
                        }
                    }
                    if (exclude.length > 0) {
                        selectParams.exclude = exclude;
                    }
                    return selectParams
                },
                processResults: function(data, params) {
                    params.page = params.page || 0;
                    var paginator = JSON.parse(data.paginator);
                    return {
                        results: data.data,
                        pagination: {
                            more: paginator.totalPage > paginator.currentPage
                        }
                    };
                }
            },
            escapeMarkup: function(markup) {
                return markup;
            }, // let our custom formatter work
            minimumInputLength: 0,
            templateResult: function(repo) {
                'use strict';
                if (repo.loading) { return repo.text; }
                return repo.name;
            },
            templateSelection: function(repo) {
                'use strict';
                return repo.name || repo.text;
            }
        });
        $("#" + AttributeLineValueId).select2({
            width: "off",
            ajax: {
                url: '/product/attributevalue/?action=search',
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        name: params.term || "", // search term
                        offset: params.page || 0,
                        limit: 5,
                    };
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    var attributeLineId = this.data("attributelineid");
                    if (attributeLineId != undefined) {
                        var attributeId = $("#" + attributeLineId).val();
                        if (attributeId == null) {
                            alert("请先选择属性,再选择属性值");
                            return;
                        } else {
                            selectParams.attributeId = attributeId;
                        }
                    }
                    if ($(this).length > 0 && $(this)[0].nodeName == "SELECT") {
                        selectParams.exclude = $(this).val();
                    }
                    return selectParams
                },
                processResults: function(data, params) {
                    params.page = params.page || 0;
                    var paginator = JSON.parse(data.paginator);
                    return {
                        results: data.data,
                        pagination: {
                            more: paginator.totalPage > paginator.currentPage
                        }
                    };
                }
            },
            escapeMarkup: function(markup) {
                return markup;
            }, // let our custom formatter work
            minimumInputLength: 0,
            templateResult: function(repo) {
                'use strict';
                if (repo.loading) { return repo.text; }
                return repo.name;
            },
            templateSelection: function(repo) {
                'use strict';
                return repo.name || repo.text;
            }
        });
    });
});
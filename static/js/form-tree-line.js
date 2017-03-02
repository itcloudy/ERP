$(function() {
    'use strict';
    var LIMIT = 5;
    var formTreeSelect2ProductAttribute = function(selector) {
        $(selector).select2({
            width: "off",
            ajax: {
                url: '/product/attribute/?action=search',
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        name: params.term || "", // search term
                        offset: (params.page || 0) * LIMIT,
                        limit: LIMIT,
                    };
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    // 过滤掉已经添加的属性
                    var existAttr = $("#product-template-attribute-body .form-tree-select-product-template-attribute");
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
                    return selectParams;
                },
                processResults: function(data, params) {
                    params.page = params.page || 0;
                    var paginator = JSON.parse(data.paginator);
                    if (data.data == undefined || data.data.length < 1) {
                        toastr.warning("没有更多可选数据", "警告");
                    }
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
                return repo.name || repo.Name;
            },
            templateSelection: function(repo) {
                'use strict';
                return repo.name || repo.Name || repo.text;
            }
        });
    };
    var formTreeSelect2ProductAttributeValues = function(selector) {
        $(selector).select2({
            width: "off",
            ajax: {
                url: '/product/attributevalue/?action=search',
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        name: params.term || "", // search term
                        offset: (params.page || 0) * LIMIT,
                        limit: LIMIT,
                    };
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    var attributeid = this.data("attributeid");
                    if (attributeid != undefined) {
                        var attributeId = $("#" + attributeid).val();
                        if (attributeId == null) {
                            // 弹框提示
                            toastr.error("请先选择<strong>属性<strong>", "错误");
                            return;
                        } else {
                            selectParams.attributeId = attributeId;
                        }
                    }
                    if ($(this).length > 0 && $(this)[0].nodeName == "SELECT") {
                        selectParams.exclude = $(this).val();
                    }
                    return selectParams;
                },
                processResults: function(data, params) {
                    params.page = params.page || 0;
                    var paginator = JSON.parse(data.paginator);
                    if (data.data == undefined || data.data.length < 1) {
                        toastr.warning("没有更多可选数据", "警告");
                    }
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
                return repo.name || repo.Name;
            },
            templateSelection: function(repo) {
                'use strict';
                return repo.name || repo.Name || repo.text;
            }
        });
    };
    // 详情页面formtree中属性和属性值页面显示
    formTreeSelect2ProductAttribute(".form-tree-select-product-template-attribute");
    formTreeSelect2ProductAttributeValues(".form-tree-select-product-template-attribute-value");
    // 产品款式增加属性
    $("#add-one-product-template-attribute").on('click', function(e) {

        var AttributeId = "AttributeId-" + e.timeStamp;
        var AttributeValueId = "AttributeValueIds-" + e.timeStamp;
        var appendHTML = "<tr data-treename='ProductAttributes' class='product-template-attribute-line form-tree-line-create'>" +
            "<td><select data-type='int' data-name='AttributeId' name='AttributeId' id='" + AttributeId + "' class='form-line-cell-create form-control form-tree-select-product-template-attribute'></select></td>" +
            "<td><select data-type='array_int' data-name='AttributeValueIds'  name='AttributeValueIds' data-attributeid='" + AttributeId + "' id='" + AttributeValueId + "' multiple='multiple' class='form-line-cell-create form-control form-tree-select-product-template-attribute-value'></select></td>" +
            "<td class='text-center'><a class='form-line-delete'><i class='fa fa-trash-o'></i></a></td>" +
            "</tr>";
        $(appendHTML).prependTo("#product-template-attribute-body");
        formTreeSelect2ProductAttribute("#" + AttributeId);
        formTreeSelect2ProductAttributeValues("#" + AttributeValueId);
    });

    //前面的代码需要重构

    // form表单中明细table行action
    $(".form-tree-line-action-remove").on('click', function(e) {
        console.log(e);
    });

});
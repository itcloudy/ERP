$(function() {
    //form中使用select2获得关联表中的数据
    $.fn.select2.defaults.set("language", "zh-CN");
    $.fn.select2.defaults.set("theme", "bootstrap");
    var LIMIT = 5;

    function formatRepo(repo) {
        'use strict';
        var name = repo.name || repo.Name;
        if (repo.loading) { return repo.text; }
        var html = "";
        html = "<p>" + name + "</p>";
        return html;
    }

    function formatRepoSelection(repo) {
        'use strict';
        var html = "";
        var name = repo.name || repo.Name;
        if (name) {
            html = "<p>" + name + "</p>";
        } else {
            html = repo.text;
        }
        return html;
    }
    var selectStaticData = function(selectClass, data) {
        'use strict';
        $(selectClass).each(function(index, el) {
            if (el.id != undefined && el.id != "") {
                var $selectNode = $("#" + el.id);
                $selectNode.select2({
                    width: "off",
                    data: data,
                    escapeMarkup: function(markup) { return markup; },
                    // minimumInputLength: 1,
                    templateResult: function(repo) {
                        if (repo.loading) { return repo.text; }
                        return repo.name;
                    },
                    templateSelection: function(repo) {
                        return repo.name;
                    }
                });
            }
        });
    };
    //selct2 Ajax 请求，现根据class选择，再根据ID绑定时间，用于后期一个页面多个相同select的情况
    var select2AjaxData = function(selectClass, ajaxUrl, tags) {
        'use strict';
        $(selectClass).each(function(index, el) {
            if (el.id != undefined && el.id != "") {
                var $selectNode = $("#" + el.id);
                nodeSelect2(el.id, ajaxUrl, tags);
            }
        });
    };
    var nodeSelect2 = function(nodeId, ajaxUrl, tags) {
        'use strict';
        $("#" + nodeId).select2({
            width: "off",
            ajax: {
                url: ajaxUrl,
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
                    if ($(this).length > 0 && $(this)[0].nodeName == "SELECT") {
                        selectParams.exclude = $(this).val();
                    }
                    return selectParams
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
            templateResult: formatRepo,
            templateSelection: formatRepoSelection
        });

    };
    select2AjaxData(".select-partner", "/partner/?action=search"); // 选择上级合伙伙伴
    select2AjaxData(".select-permission", "/permission/?action=search"); // 选择权限
    select2AjaxData(".select-role", "/role/?action=search"); // 选择角色
    select2AjaxData(".select-source", "/source/?action=search"); // 选择菜单
    select2AjaxData(".select-menu", "/menu/?action=search"); // 系统资源
    select2AjaxData(".select-user", "/user/?action=search"); // 选择用户
    select2AjaxData(".select-company", "/company/?action=search"); // 选择公司
    select2AjaxData(".select-department", "/department/?action=search"); // 选择部门
    select2AjaxData(".select-position", "/position/?action=search"); // 选择职位
    select2AjaxData(".select-team", "/team/?action=search", true); // 选择团队
    select2AjaxData(".select-product-counter", "/product/counter/?action=search"); // 选择产品柜台
    select2AjaxData(".select-product-category", "/product/category/?action=search"); // 选择产品类别;
    select2AjaxData(".select-product-attribute", '/product/attribute/?action=search'); // 选择属性
    select2AjaxData(".select-product-attribute-value", '/product/attributevalue/?action=search'); // 选择属性值
    selectStaticData(".select-product-type", [{ id: "stock", name: '库存商品' }, { id: "consume", name: '消耗品' }, { id: "service", name: '服务' }]); // 产品类型
    select2AjaxData(".select-product-uom", "/product/uom/?action=search"); // 选择产品单位
    select2AjaxData(".select-product-uom-category", "/product/uomcateg/?action=search"); //计量单位类别
    selectStaticData(".select-product-uom-category-type", [{ id: 1, name: '小于参考计量单位' }, { id: 2, name: '参考计量单位' }, { id: 3, name: '大于参考计量单位' }]); // 产品类型
    //地址选择
    var addressSelectData = function(selectClass, ajaxUrl) {
        $(selectClass).select2({
            width: "off",
            ajax: {
                url: ajaxUrl,
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        Name: params.term || "", // search term
                        DefaultCode: params.term || "",
                        offset: (params.page || 0) * LIMIT,
                        limit: LIMIT,
                    };
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    var selectId = this.attr("id");
                    if (selectId == "district") {
                        var city = $("#city");
                        if (city.length < 1) {
                            toastr.error("没有<strong>城市</strong>选项", "错误");
                            return;
                        } else {
                            city = city.val();
                            if (city == null || city == undefined) {
                                toastr.error("请按照<strong>国家->省份->城市->区县</strong>的顺序选择", "错误");
                                return;
                            } else {
                                selectParams.CityID = parseInt(city);
                            }
                        }
                    } else if (selectId == "city") {
                        var province = $("#province");
                        if (province.length < 1) {
                            toastr.error("没有<strong>省份</strong>选项", "错误");
                            return;
                        } else {
                            province = province.val();
                            if (province == null || province == undefined) {
                                toastr.error("请按照<strong>国家->省份->城市->区县</strong>的顺序选择", "错误");
                                return;
                            } else {
                                selectParams.ProvinceID = parseInt(province);
                            }
                        }
                    } else if (selectId == "province") {
                        var country = $("#country");
                        if (country.length < 1) {
                            toastr.error("没有<strong>国家</strong>选项", "错误");
                            return;
                        } else {
                            country = country.val();
                            if (country == null || country == undefined) {
                                toastr.error("请按照<strong>国家->省份->城市->区县</strong>的顺序选择", "错误");
                                return;
                            } else {
                                selectParams.CountryID = parseInt(country);
                            }
                        }
                    }
                    return selectParams
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
                var html = "";
                html = "<p>" + repo.Name + "</p>";
                return html;
            },
            templateSelection: function(repo) {
                'use strict';
                var html = "";
                if (repo.Name != undefined) {
                    html = "<p>" + repo.Name + "</p>";
                } else {
                    html = repo.text;
                }

                return html;
            }
        });
    };
    addressSelectData(".select-address-country", "/address/country/?action=search"); // 选择国家
    addressSelectData(".select-address-province", "/address/province/?action=search"); // 选择省份
    addressSelectData(".select-address-city", "/address/city/?action=search"); // 选择城市
    addressSelectData(".select-address-district", "/address/district/?action=search"); // 选择地区
    // 根据款式创建产品，款式修改后，需要同时更新产品的类别，销售和采购单位
    $(".select-product-template").select2({
        width: "off",
        ajax: {
            url: "/product/template/?action=search",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params) {
                var selectParams = {
                    Name: params.term || "", // search term
                    DefaultCode: params.term || "",
                    offset: (params.page || 0) * LIMIT,
                    limit: LIMIT,
                };
                var xsrf = $("input[name ='_xsrf']");
                if (xsrf.length > 0) {
                    selectParams._xsrf = xsrf[0].value;
                }
                if ($(this).length > 0 && $(this)[0].nodeName == "SELECT") {
                    selectParams.exclude = $(this).val();
                }
                return selectParams
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
            var html = "";
            html = "<p><strong>款式编码:&nbsp</strong>" + repo.DefaultCode + "</p>";
            html += "<p><strong>款式名称:&nbsp</strong>" + repo.Name + "</p>";
            return html;
        },
        templateSelection: function(repo) {
            'use strict';
            var html = "";
            if (repo.Name != undefined) {
                html = "<span style='color:#337ab7;'>[" + repo.DefaultCode + "]</span>" + repo.Name;
            } else {
                var $option = $(repo.element);
                var defaultCode = $option.data("defaultcode");
                if (defaultCode != undefined) {
                    html = "<span style='color:#337ab7;'>[" + defaultCode + "]</span>" + repo.text;
                } else {
                    html = repo.text;
                }

            }
            return html;
        }
    }).on("change", function(e) {
        var productTempId = parseInt(e.currentTarget.value);
        $.ajax({
            type: 'POST',
            url: "/product/template/?action=search",
            data: (function() {
                var params = { Id: productTempId };
                var xsrf = $("input[name ='_xsrf']");
                if (xsrf.length > 0) {
                    params._xsrf = xsrf[0].value;
                }
                return params;
            })(),
            success: function(result) {
                if (result.data && result.data.length > 0) {
                    var Pdata = result.data[0];
                    $("#name").val(Pdata.Name);
                    $("#product-attributevalues").empty();
                    $("#category").empty().append("<option value='" + Pdata.Category.id + "' selected='selected'>" + Pdata.Category.name + "</option>");
                    $("#firstSaleUom").empty().append("<option value='" + Pdata.FirstSaleUom.id + "' selected='selected'>" + Pdata.FirstSaleUom.name + "</option>");
                    if (Pdata.SecondSaleUom != undefined) {
                        $("#secondSaleUom").empty().append("<option value='" + Pdata.SecondSaleUom.id + "' selected='selected'>" + Pdata.SecondSaleUom.name + "</option>");
                    }
                    $("#firstPurchaseUom").empty().append("<option value='" + Pdata.FirstPurchaseUom.id + "' selected='selected'>" + Pdata.FirstPurchaseUom.name + "</option>");
                    if (Pdata.ProductCounter != undefined) {
                        $("#productCounter").empty().append("<option value='" + Pdata.ProductCounter.id + "' selected='selected'>" + Pdata.ProductCounter.name + "</option>");
                    }
                    if (Pdata.SecondPurchaseUom != undefined) {
                        $("#secondPurchaseUom").empty().append("<option value='" + Pdata.SecondPurchaseUom.id + "' selected='selected'>" + Pdata.SecondPurchaseUom.name + "</option>");
                    }
                    // 重置表单验证
                    $("#productProductForm").data('bootstrapValidator').resetForm();
                }
            },
            dataType: "json"
        });
    });
    var selectProductProductAttributeValue = function(selector) {
        $(selector).select2({
            width: "off",
            ajax: {
                url: '/product/attributeline/?action=search',
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params) {
                    var selectParams = {
                        name: params.term || "", // search term
                        offset: (params.page || 0) * LIMIT,
                        limit: 5,
                        productAttrs: true
                    };
                    var tmpId = $("#ProductTemplateID");
                    if (tmpId.length < 1) {
                        toastr.error("请先选择<strong>款式</strong>", "错误");
                        return;
                    } else {
                        tmpId = tmpId.val();
                        if (tmpId == null || tmpId == undefined) {
                            toastr.error("请先选择<strong>款式</strong>", "错误");
                            return;
                        } else {
                            selectParams.tmpId = parseInt(tmpId);
                        }
                    }
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf.length > 0) {
                        selectParams._xsrf = xsrf[0].value;
                    }
                    var attributeid = this.data("attributeid");
                    if (attributeid != undefined) {
                        var attributeId = $("#" + attributeid).val();
                        if (attributeId == null) {
                            // 弹框提示
                            toastr.error("请先选择<strong>属性</strong>", "错误");
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
                    var data = data.data;
                    var results = [];
                    if (data && data.length > 0) {
                        for (var i = 0, len1 = data.length; i < len1; i++) {
                            var attributeValueArrs = data[i].attributeValueArrs;
                            for (var j = 0, len2 = attributeValueArrs.length; j < len2; j++) {
                                results.push({
                                    id: attributeValueArrs[j].id,
                                    name: attributeValueArrs[j].name
                                })
                            }
                        }
                    }
                    return {
                        results: results,
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
    };
    selectProductProductAttributeValue(".select-product-product-attribute-value");
});
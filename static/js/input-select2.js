//form中使用select2获得关联表中的数据
$.fn.select2.defaults.set("language", "zh-CN");
$.fn.select2.defaults.set("theme", "bootstrap");
var LIMIT = 5;

function formatRepo(repo) {
    'use strict';
    if (repo.loading) { return repo.text; }
    return repo.name;
}

function formatRepoSelection(repo) {
    'use strict';
    return repo.name || repo.text;
}
var selectStaticData = function(selectClass, data) {
    'use strict';
    $(selectClass).each(function(index, el) {
        if (el.id != undefined && el.id != "") {
            var $selectNode = $("#" + el.id);
            $selectNode.select2({
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
            Nodeselect2(el.id, ajaxUrl, tags);
        }
    });
};
var Nodeselect2 = function(nodeId, ajaxUrl, tags) {
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
                    offset: params.page || 0,
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
                return {
                    results: data.data,
                    pagination: {
                        more: paginator.totalPage > paginator.currentPage
                    }
                };
            },
        },
        escapeMarkup: function(markup) {
            return markup;
        }, // let our custom formatter work
        minimumInputLength: 0,
        templateResult: formatRepo,
        templateSelection: formatRepoSelection
    });

};
select2AjaxData(".select-department", "/department/?action=search"); // 选择部门
select2AjaxData(".select-position", "/position/?action=search"); // 选择职位
select2AjaxData(".select-group", "/group/?action=search", true); // 选择分组
select2AjaxData(".select-product-category", "/product/category/?action=search"); // 选择产品类别;
select2AjaxData(".select-product-attribute", '/product/attribute/?action=search'); // 选择属性
// selectStaticData(".select-product-type", [{ id: 1, name: '库存商品' }, { id: 2, name: '消耗品' }, { id: 3, name: '服务' }]); // 产品类型
select2AjaxData(".select-product-uom", "/product/uom/?action=search"); // 选择产品单位
select2AjaxData(".select-product-uom-category", "/product/uomcateg/?action=search"); //计量单位类别
selectStaticData(".select-product-uom-category-type", [{ id: 1, name: '小于参考计量单位' }, { id: 2, name: '参考计量单位' }, { id: 3, name: '大于参考计量单位' }]); // 产品类型
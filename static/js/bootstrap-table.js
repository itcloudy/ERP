//tabel视图中使用bootstrap-table来显示数据
$.extend($.fn.bootstrapTable.defaults, {
    method: "post",
    dataType: "json",
    locale: "zh-CN",
    contentType: "application/x-www-form-urlencoded",
    sidePagination: "server",
    stickyHeader: true, //表头固定
    stickyHeaderOffsetY: (function() {
        'use strict';
        var stickyHeaderOffsetY = 0;
        if ($('.navbar-fixed-top').css('height')) {
            stickyHeaderOffsetY = +$('.navbar-fixed-top').css('height').replace('px', '');
        }
        if ($('.navbar-fixed-top').css('margin-bottom')) {
            stickyHeaderOffsetY += +$('.navbar-fixed-top').css('margin-bottom').replace('px', '');
        }
        return stickyHeaderOffsetY + 'px';
    })(), //设置偏移量
    dataField: "data",
    pagination: true,
    pageNumber: 1,
    pageSize: 20,
    pageList: [10, 25, 50, 100, 500, 1000],
    // onClickRow: function(row, $element) {
    //     //$element是当前tr的jquery对象
    //     $element.css("background-color", "green");
    // },//单击row事件
});
var displayTable = function(selectId, ajaxUrl, columns, onExpandRow) {
    var $tableNode = $(selectId);
    var options = {
        url: ajaxUrl,
        queryParams: function(params) {
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action = 'table';
            var filterCond = $(".list-info-table .form-control");
            var filter = {};
            //获得过滤条件
            if (filterCond.length > 0) {
                filterCond.each(function() {
                    if (this.type == 'text') {
                        if (this.value != "") {
                            filter[this.name] = this.value;
                        }
                    }
                });
            }
            params.filter = JSON.stringify(filter);
            return params;
        },
        columns: columns
    }
    if (onExpandRow != undefined) {
        options.detailView = true;
        options.onExpandRow = onExpandRow;
    }
    $tableNode.bootstrapTable(options);
};
//用户表
displayTable("#table-user", "/user/", [
    { title: "全选", field: 'Id', checkbox: true, align: "center", valign: "middle" },
    { title: "用户名", field: 'Name', sortable: true, order: "desc" },
    { title: "中文名称", field: 'NameZh', sortable: true, order: "desc" },
    { title: "部门", field: 'Department', sortable: true, order: "desc", filter: { type: "select", data: [] } },
    { title: "职位", field: 'Position', sortable: true, order: "desc", filter: { type: "select", data: [] } },
    { title: "邮箱", field: 'Email', sortable: true, order: "desc" },
    { title: "手机号码", field: 'Mobile', sortable: true, order: "desc" },
    { title: "座机", field: 'Tel', sortable: true, order: "desc" },
    { title: "QQ", field: 'Qq', sortable: true, order: "desc" },
    { title: "微信", field: 'Wechat', sortable: true, order: "desc" },
    {
        title: "管理员",
        field: 'IsAdmin',
        sortable: true,
        order: "desc",
        align: "center",

        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsAdmin) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    {
        title: "操作",
        align: "left",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = '';
            var url = "/user/";
            if (row.Active) {
                html += "<a href='" + url + row.ID + "?action=invalid' class='btn btn-xs btn-default table-action '><i class='fa fa-close'>&nbsp无效</i></a>";
            } else {
                html += "<a href='" + url + row.ID + "?action=active' class='btn btn-xs btn-default table-action '><i class='fa fa-check'>&nbsp有效</i></a>";
            }
            html += "<a href='" + url + row.ID + "?action=edit' class='btn btn-xs btn-default table-action '><i class='fa fa-pencil'>&nbsp编辑</i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='btn btn-xs btn-default table-action '><i class='fa fa-external-link'>&nbsp详情</i></a>";
            return html;
        }
    }
], function(index, row, $detail) {
    var params = (function() {
        var params = {};
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf != undefined) {
            params._xsrf = xsrf[0].value;
        }
        params.action = 'table';
        params.offset = 0;
        params.limit = 5;
        return params;
    })();
    $.ajax({
        url: "/user/",
        dataType: "json",
        type: "POST",
        async: false,
        data: params,
        success: function(data) {
            html = "ok";
            $detail.html(data.total);
        },
        error: function(error) {

            html = error;
            $detail.html(html);
        }
    });
});
//权限
displayTable("#table-group", "/group/", [
    { title: "全选", field: 'Id', checkbox: true, align: "center", valign: "middle" },
    { title: "权限组名", field: 'name', sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'active',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.active) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    { title: "定位", field: 'location', sortable: true, order: "desc" },
    { title: "描述", field: 'description', sortable: true },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/group/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//登录记录表
displayTable("#table-record", "/record/", [
    { title: "全选", field: 'Id', checkbox: true, align: "center", valign: "middle" },
    { title: "用户名", field: 'Name', sortable: true, order: "desc" },
    { title: "邮箱", field: 'Email', sortable: true, order: "desc" },
    { title: "手机号码", field: 'Mobile', sortable: true, order: "desc" },
    { title: "开始时间", field: 'CreateDate', sortable: true, order: "desc" },
    {
        title: "结束时间",
        field: 'Logout',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Logout == "0001-01-01 00:00:00") {
                html = "<small>非正常退出</small>";
            } else {
                html = row.Logout;
            }
            return html;
        }
    },
    { title: "IP地址", field: 'Ip', sortable: true, order: "desc" },
    { title: "用户代理", field: "UserAgent" }

]);
displayTable("#table-sequence", "/sequence", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "序号名称", field: 'Name', sortable: true, order: "desc" },
    { title: "表名称", field: "StructName", sortable: true, order: "desc" },
    { title: "前缀", field: "Prefix", sortable: true, order: "desc" },
    { title: "数字位数", field: "Padding", sortable: true, order: "desc" },
    { title: "当前序号", field: "Current", sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    }, {
        title: "为默认",
        field: 'IsDefault',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsDefault) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },

    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/sequence/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//国家表
displayTable("#table-country", "/address/country/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "国家", field: 'Name', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/country/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//省份表
displayTable("#table-province", "/address/province/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "省份", field: 'Name', sortable: true, order: "desc" },
    { title: "国家", field: 'Country', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/province/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//城市表
displayTable("#table-city", "/address/city/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "城市", field: 'Name', sortable: true, order: "desc" },
    { title: "省份", field: 'Province', sortable: true, order: "desc" },
    { title: "国家", field: 'Country', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/city/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//区县表
displayTable("#table-district", "/address/district/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "地区", field: 'Name', sortable: true, order: "desc" },
    { title: "城市", field: 'City', sortable: true, order: "desc" },
    { title: "省份", field: 'Province', sortable: true, order: "desc" },
    { title: "国家", field: 'Country', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/district/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品属性
displayTable("#table-product-attribute", "/product/attribute/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "属性名", field: 'name', sortable: true, order: "desc" },
    { title: "属性编码", field: 'code', sortable: true, order: "desc" },
    { title: "属性序号", field: 'sequence', sortable: true, order: "desc" },
    {
        title: "属性值",
        field: 'childs',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.values;
            var html = "";
            var url = "/product/attributevalue/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + "</a>";
            }
            return html;
        }
    },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/attribute/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品类别
displayTable("#table-product-category", "/product/category/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "类别名", field: 'name', sortable: true, order: "desc" },
    { title: "上级", field: 'parent', sortable: true, order: "desc" },
    { title: "上级路径", field: 'path', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/category/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品款式
displayTable("#table-product-template", "/product/template/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "款式编码", field: 'DefaultCode', sortable: true, order: "desc" },
    { title: "款式名称", field: 'Name', sortable: true, order: "desc" },
    { title: "款式类别", field: 'Category', sortable: true, order: "desc" },
    { title: "规格数量", field: 'VariantCount', sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    {
        title: "可销售",
        field: 'SaleOk',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.SaleOk) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    {
        title: "规格创建方式",
        field: 'ProductMethod',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.ProductMethod == "auto") {
                html = '自动';
            } else if (row.ProductMethod == "hand") {
                html = '手动';
            } else {
                html = '-';
            }
            return html;
        }
    },
    {
        title: "款式类型",
        field: 'ProductType',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.ProductType == "stock") {
                html = '库存商品';
            } else if (row.ProductType == "consume") {
                html = '消耗品';
            } else if (row.ProductType == "service") {
                html = '服务';
            } else {
                html = '-';
            }
            return html;
        }
    },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/template/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }

]);

//产品规格
displayTable("#table-product-product", "/product/product/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "规格编码", field: 'defaultCode', sortable: true, order: "desc" },
    { title: "规格名次", field: 'name', sortable: true, order: "desc" },
    { title: "规格类别", field: 'category', sortable: true, order: "desc" },
    { title: "产品款式", field: 'parent', sortable: true, order: "desc" },
    { title: "规格属性", field: 'attributes', align: "center", sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/product/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品属性值
displayTable("#table-product-attributevalue", "/product/attributevalue/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "属性", field: 'attribute', sortable: true, order: "desc" },
    { title: "属性值", field: 'name', align: "center", sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/attributevalue/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品单位类别
displayTable("#table-product-uom-categ", "/product/uomcateg/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "计量单位类别", field: 'name', sortable: true, order: "desc" },
    {
        title: "计量单位",
        field: 'uoms',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.uoms;
            var html = "";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='/product/uom/" + key + "?action=detail'>" + datas[key] + "</a>";
            }
            return html;
        }

    },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/uomcateg/";
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品单位
displayTable("#table-product-uom", "/product/uom/", [
    { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
    { title: "计量单位类别", field: 'category', sortable: true, order: "desc" },
    { title: "计量单位", field: 'name', align: "center", sortable: true, order: "desc" },
    { title: "符号", field: 'symbol', align: "center", },
    { title: "类型", field: 'type', align: "center", sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'active',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.active) {
                html = '<i class="fa fa-check"></i>';
            } else {
                html = '<i class="fa fa-remove"></i>';
            }
            return html;
        }
    },
    { title: "比率", field: 'factor', align: "center", sortable: true, order: "desc" },
    { title: "更大比率", field: 'factorInv', align: "center", sortable: true, order: "desc" },
    { title: "舍入精度", field: 'rounding', align: "center", sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {

            var html = "";
            var url = "/product/uom/";
            if (row.active) {
                html += "<a href='" + url + row.ID + "?action=invalid' class='table-action btn btn-xs btn-default'>无效&nbsp<i class='fa fa-close'></i></a>";
            } else {
                html += "<a href='" + url + row.ID + "?action=active' class='table-action btn btn-xs btn-default'>有效&nbsp<i class='fa fa-check'></i></a>";
            }
            html += "<a href='" + url + row.ID + "?action=edit' class='table-action btn btn-xs btn-default'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.ID + "?action=detail' class='table-action btn btn-xs btn-default'>详情&nbsp<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//鼠标进入展示详情，离开隐藏详情
// $('#table-user').on('mouseenter mouseleave', 'tbody>tr',
//     function(e) {
//         $(this).find(".detail-icon").trigger("click");
//     }
// );
$(".list-info-table .form-control").change(function(e) {
    $(".table-diplay-info").bootstrapTable('refresh');
});
$("#clearListSearchCond-table ").click(function() {
    $(".table-diplay-info").bootstrapTable('refresh');
});
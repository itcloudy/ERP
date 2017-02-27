//鼠标进入展示详情，离开隐藏详情
// $('#table-user').on('mouseenter mouseleave', 'tbody>tr',
//     function(e) {
//         $(this).find(".detail-icon").trigger("click");
//     }
// );
$(".list-info-table .form-control").change(function(e) {
    $(".table-diplay-info").bootstrapTable('refresh');
});
$("#clearListSearchCond-table").click(function() {
    $(".table-diplay-info").bootstrapTable('refresh');
});
$("#listViewSearch input").on("change", function(e) { console.log(e); });
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
    // }, //单击row事件
    onCheck: function(row, $el) {
        $($el[0].parentNode.parentNode).addClass("danger");
    },
    onUncheck: function(row, $el) {
        $($el[0].parentNode.parentNode).removeClass("danger");
    },
    onCheckAll: function(rows) {
        $("#display-table tbody>tr").addClass("danger");
    },
    onUncheckAll: function(rows) {
        $("#display-table tbody>tr").removeClass("danger");
    }
});

//根据数据类型获得正确的数据,默认string
function getCurrentDataType(val, dataType) {
    if (dataType == "" || dataType === undefined || dataType === null) {
        dataType = "string";
    }
    switch (dataType) {
        case "int": // 整形
            val = parseInt(val);
            break;
        case "float": // 浮点型
            val = parseFloat(val);
            break;
        case "array_int": // 整形数组
            var a_arr = [];
            for (var a_i = 0, a_l = val.length; a_i < a_l; a_i++) {
                a_arr.push(parseInt(val[a_i]));
            }
            val = a_arr;
            break;
        case "arrar_float": //  浮点型数组
            var a_arr = [];
            for (var a_i = 0, a_l = val.length; a_i < a_l; a_i++) {
                a_arr.push(parseFloat(val[a_i]));
            }
            val = a_arr;
            break;
    }
    return val
};

function defaultQueryParams(params) {
    var xsrf = $("input[name ='_xsrf']");
    if (xsrf != undefined) {
        params._xsrf = xsrf[0].value;
    }
    params.action = 'table';
    var filterCond = $("#listViewSearch .filter-condition");
    var filter = {};
    // console.log(filterCond);
    //获得过滤条件
    for (var i = 0, len = filterCond.length; i < len; i++) {
        var self = filterCond[i];
        // 处理radio数据
        if (self.type == "radio") {
            if ($(self).data("type") == "string") {
                var nodeName = $("input[name ='" + self.name + "']:checked");
                if (nodeName != undefined) {
                    filter[self.name] = nodeName.val();
                }
            } else {
                console.log("data  type is not string");
            }
        } else if (self.type == "checkbox") {
            if (self.checked) {
                filter[self.name] = true;
            } else {
                filter[self.name] = false;
            }
        } else {
            var val = $(self).val();
            if (val != "") {
                // 若为null跳出此次循环
                if (val === null) {
                    continue;
                }
                filter[self.name] = getCurrentDataType(val, $(self).data("type"))
            }
        }
    }
    params.filter = JSON.stringify(filter);
    return params;
};
var displayTable = function(selectId, ajaxUrl, columns, bootstrapTableFunctionDict) {
    var $tableNode = $(selectId);
    var queryParams = defaultQueryParams;
    var onExpandRow = undefined;
    var onPostBody = undefined;
    if (bootstrapTableFunctionDict != undefined) {
        if (bootstrapTableFunctionDict.onExpandRow != undefined) {
            onExpandRow = bootstrapTableFunctionDict.onExpandRow;
        }
        if (bootstrapTableFunctionDict.onPostBody != undefined) {
            onPostBody = bootstrapTableFunctionDict.onPostBody;
        }
        if (bootstrapTableFunctionDict.queryParams) {
            queryParams = bootstrapTableFunctionDict.queryParams;
        }
    }


    var options = {
        url: ajaxUrl,
        queryParams: queryParams,
        columns: columns
    }
    if (onExpandRow != undefined) {
        options.detailView = true;
        options.onExpandRow = onExpandRow;
    }
    if (onPostBody != undefined) {
        options.onPostBody = onPostBody;
    }

    $tableNode.bootstrapTable(options);
    // 选中行颜色变化
    // $("#display-table .bs-checkbox").on('click', function(e) {
    //     console.log(e);
    // });
};
//用户表
displayTable("#table-user", "/user/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "用户名", field: 'Name', sortable: true, order: "desc" },
    { title: "中文名称", field: 'NameZh', sortable: true, order: "desc" },
    { title: "所属公司", field: 'Company', sortable: true, order: "desc" },
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
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        class: "data-active",
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
            html += "<a href='" + url + row.id + "?action=edit' class='btn btn-xs btn-default table-action '><i class='fa fa-pencil'>编辑</i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='btn btn-xs btn-default table-action '><i class='fa fa-external-link'>详情</i></a>";
            return html;
        }
    }
], {
    onExpandRow: function(index, row, $detail) {
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
    }
});
// 公司
displayTable("#table-company", '/company/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "公司名称", field: 'Name', sortable: true, order: "desc" },
    { title: "公司编码", field: 'Code', sortable: true, order: "desc" },
    { title: "母公司", field: 'Parent', sortable: true, order: "desc" },
    { title: "公司地址", field: 'Address' },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/company/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
// 部门
displayTable("#table-department", '/department/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "部门名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "负责人",
        field: 'Leader',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.Leader) {
                html = row.Leader.name + "<a class='pull-right' href='/user/" + row.Leader.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
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
            var url = "/department/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
// 团队
displayTable("#table-team", '/team/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "团队名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "负责人",
        field: 'Leader',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Leader) {
                html = row.Leader.name + "<a class='pull-right' href='/user/" + row.Leader.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    {
        title: "团队成员",
        field: 'Members',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var datas = row.Members;
            var html = "";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='/product/uom/" + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
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
            var url = "/team/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
// 职位
displayTable("#table-position", '/position/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "职位名称", field: 'Name', sortable: true, order: "desc" },
    { title: "职位描述", field: 'Description' },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/position/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }

]);
//系统资源
displayTable("#table-source", "/source/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "资源名称", field: 'Name', sortable: true, order: "desc" },
    { title: "Model名称", field: 'ModelName', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/source/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//角色
displayTable("#table-role", "/role/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "角色名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "角色用户",
        field: 'Users',
        formatter: function cellStyle(value, row, index) {
            var datas = row.Users;
            var html = "";
            var url = "/role/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
            }
            return html;
        }
    },
    {
        title: "权限列表",
        field: 'Permissions',
        formatter: function cellStyle(value, row, index) {
            var datas = row.Permissions;
            var html = "";
            var url = "/permission/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
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
            var url = "/role/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//权限
displayTable("#table-permission", "/permission/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "权限名称", field: 'Name', sortable: true, order: "desc" },
    { title: "资源名称", field: 'Source', sortable: true, order: "desc" },
    {
        title: "权限类型",
        field: 'Relation',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.ProductType == "owner") {
                html = '私有权限';
            } else if (row.ProductType == "role") {
                html = '角色权限';
            } else {
                html = '-';
            }
            return html;
        }
    },
    {
        title: "创建权限",
        field: 'PermCreate',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PermCreate) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "创建权限",
        field: 'PermCreate',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PermCreate) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "查询权限",
        field: 'PermRead',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PermRead) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "修改权限",
        field: 'PermWrite',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PermWrite) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "删除权限",
        field: 'PermDelete',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PermCreate) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
            var url = "/permission/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
displayTable("#table-menu", '/menu/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "菜单名", field: 'Name', sortable: true, order: "desc" },
    { title: "菜单唯一标识", field: 'Identity', sortable: true, order: "desc" },
    {
        title: "菜单可见角色",
        field: 'Roles',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.Roles;
            var html = "";
            var url = "/role/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
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
            var url = "/menu/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//登录记录表
displayTable("#table-record", "/record/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
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
// 序号管理
displayTable("#table-sequence", "/sequence", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "所属公司", field: 'Company', sortable: true, order: "desc" },
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
        class: "data-active",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//国家表
displayTable("#table-address-country", "/address/country/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "国家", field: 'Name', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/country/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//省份表
displayTable("#table-address-province", "/address/province/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "省份", field: 'Name', sortable: true, order: "desc" },
    { title: "国家", field: 'Country', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/address/province/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//城市表
displayTable("#table-address-city", "/address/city/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
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
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//区县表
displayTable("#table-address-district", "/address/district/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "区县", field: 'Name', sortable: true, order: "desc" },
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
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
// 合作伙伴管理
displayTable("#table-partner", '/partner/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "是公司",
        field: 'IsCompany',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsCompany) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "是客户",
        field: 'IsCustomer',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsCustomer) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "是供应商",
        field: 'IsSupplier',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsSupplier) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    { title: "邮箱", field: 'Email', sortable: true, order: "desc" },
    { title: "手机号码", field: 'Mobile', sortable: true, order: "desc" },
    { title: "电话号码", field: 'Tel', sortable: true, order: "desc" },
    { title: "QQ", field: 'Qq', sortable: true, order: "desc" },
    { title: "微信", field: 'WeChat', sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/partner/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品属性
displayTable("#table-product-attribute", "/product/attribute/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "属性名", field: 'Name', sortable: true, order: "desc" },
    { title: "属性编码", field: 'Code', sortable: true, order: "desc" },
    { title: "属性序号", field: 'Sequence', sortable: true, order: "desc" },
    {
        title: "属性值",
        field: 'childs',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.values;
            var html = "";
            var url = "/product/attributevalue/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
            }
            return html;
        }
    },
    { title: "产品款式数量", field: 'TemplatesCount', align: "center", sortable: true, order: "desc" },
    { title: "产品规格数量", field: 'ProductsCount', align: "center", sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/attribute/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
// 产品柜台
displayTable("#table-sale-counter", "/sale/counter/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "柜台名称", field: 'Name', sortable: true, order: "desc" },
    { title: "产品款式数量", field: 'TemplatesCount', align: "center" },
    { title: "产品规格数量", field: 'ProductsCount', align: "center" },
    { title: "柜台描述", field: 'Description' },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/sale/counter/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品类别
displayTable("#table-product-category", "/product/category/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
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
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品款式
displayTable("#table-product-template", "/product/template/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "款式编码", field: 'DefaultCode', sortable: true, order: "desc" },
    { title: "款式名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "款式类别",
        field: 'Category',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            return row.Category.name;
        }
    },
    { title: "规格数量", field: 'VariantCount', align: "center", sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        order: "desc",
        align: "center",
        class: "data-active",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
        title: "产品款式类型",
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
        title: "第一销售单位",
        field: 'FirstSaleUom',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            return row.FirstSaleUom.name;
        }
    },
    {
        title: "第一采购单位",
        field: 'FirstPurchaseUom',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            return row.FirstPurchaseUom.name;
        }
    },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/template/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }

]);
//产品款式属性明细
displayTable("#table-product-attribute-line", "/product/attributeline/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "款式编码", field: 'ProductTemplate.DefaultCode', sortable: true, order: "desc" },
    { title: "产品款式", field: 'ProductTemplate', sortable: true, order: "desc" },
    { title: "属性", field: 'Attribute', align: "center", sortable: true, order: "desc" },
    {
        title: "属性值",
        field: 'Attribute',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.attributeValueArrs.length > 0) {
                var attributeValueArrs = row.attributeValueArrs;
                for (var i = 0, len = attributeValueArrs.length; i < len; i++) {
                    html += "<a  class='display-block label label-success' href='/product/attributevalue/" + attributeValueArrs[i].id + "?action=detail'>" + attributeValueArrs[i].name + "</a>";
                }
            }
            return html;
        }
    },


]);

//产品规格
displayTable("#table-product-product", "/product/product/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "规格编码", field: 'DefaultCode', sortable: true, order: "desc" },
    { title: "规格名称", field: 'Name', sortable: true, order: "desc" },
    {
        title: "规格类别",
        field: 'Category',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            return row.Category.name;
        }
    },
    {
        title: "产品款式",
        field: 'ProductTemplate',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = row.ProductTemplate.name + "<a class='pull-right' href='/product/template/" + row.ProductTemplate.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            return html;
        }
    },
    {
        title: "规格属性",
        field: 'AttributeValues',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.AttributeValues;
            var html = "";
            var url = "/product/attributevalue/";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
            }
            return html;
        }
    },
    {
        title: "有效",
        field: 'Active',
        sortable: true,
        order: "desc",
        class: "data-active",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "产品规格类型",
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
        title: "第一销售单位",
        field: 'FirstSaleUom',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            return row.FirstSaleUom.name;
        }
    },
    {
        title: "第一采购单位",
        field: 'FirstPurchaseUom',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            return row.FirstPurchaseUom.name;
        }
    },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/product/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);

//产品属性值
displayTable("#table-product-attributevalue", "/product/attributevalue/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "属性", field: 'Attribute', sortable: true, order: "desc" },
    { title: "属性值", field: 'Name', align: "center", sortable: true, order: "desc" },
    { title: "产品规格数量", field: 'ProductsCount', align: "center", sortable: true, order: "desc" },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/product/attributevalue/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品单位类别
displayTable("#table-product-uom-categ", "/product/uomcateg/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "计量单位类别", field: 'name', sortable: true, order: "desc" },
    {
        title: "计量单位",
        field: 'uoms',
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var datas = row.uoms;
            var html = "";
            for (key in datas) {
                html += "<a  class='display-block label label-success' href='/product/uom/" + key + "?action=detail'>" + datas[key] + '<span style="display:none;">;<span>' + "</a>";
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
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
//产品单位
displayTable("#table-product-uom", "/product/uom/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
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
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
                html += "<a href='" + url + row.id + "?action=invalid' class='table-action btn btn-xs btn-default'>无效<i class='fa fa-close'></i></a>";
            } else {
                html += "<a href='" + url + row.id + "?action=active' class='table-action btn btn-xs btn-default'>有效<i class='fa fa-check'></i></a>";
            }
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
displayTable("#table-stock-warehouse", "/stock/warehouse/", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    {
        title: "所属公司",
        field: 'Company',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Company) {
                html = row.Company.name + "<a class='pull-right' href='/company/" + row.Company.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    { title: "仓库名称", field: 'Name', sortable: true, order: "desc" },
    { title: "发货库位", field: 'Location', sortable: true, order: "desc" },
    { title: "仓库编码", field: 'Code', align: "center", sortable: true, order: "desc" },
    { title: "仓库地址", field: 'Address' },
    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var url = "/stock/warehouse/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
displayTable("#table-stock-picking-type", '/stock/picking/type/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "库位类型", field: 'Name', align: "center", sortable: true, order: "desc" },
    {
        title: "所属仓库",
        field: 'WareHouse',
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {

            var html = "";
            var url = "/stock/warehouse/";
            if (row.WareHouse) {
                html += row.WareHouse.name + "<a href='" + url + row.WareHouse.id + "?action=detail' class='pull-right'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    {
        title: "移库类型",
        field: 'Code',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Code == "outgoing") {
                html = '出库';
            } else if (row.Code == "incoming") {
                html = '入库';
            } else if (row.Code == "internal") {
                html = '内部调拨';
            } else {
                html = '-';
            }
            return html;
        }
    },
    {
        title: "上步流程",
        field: 'PrevStep',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.PrevStep) {
                html = row.PrevStep.name + "<a class='pull-right' href='/stock/picking/type/" + row.PrevStep.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    {
        title: "下步流程",
        field: 'NextStep',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.NextStep) {
                html = row.NextStep.name + "<a class='pull-right' href='/stock/picking/type/" + row.NextStep.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    {
        title: "流程开始",
        field: 'IsStart',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsStart) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "流程结束",
        field: 'IsEnd',
        sortable: true,
        order: "desc",
        align: "center",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.IsEnd) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
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
            var url = "/stock/picking/type/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }

]);
displayTable("#table-stock-location", '/stock/location/', [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    {
        title: "所属公司",
        field: 'Company',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Company) {
                html = row.Company.name + "<a class='pull-right' href='/company/" + row.Company.id + "?action=detail'><i class='fa fa-external-link'></i></a>";
            }
            return html;
        }
    },
    { title: "库位名称", field: 'Name', align: "center", sortable: true, order: "desc" },
    {
        title: "库位类型",
        field: 'Usage',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            var usage = row.Usage;
            if ("supplier" == usage) {
                html = "供应商库位";
            } else if ("view" == usage) {
                html = "视图";
            } else if ("internal" == usage) {
                html = "内部库位";
            } else if ("customer" == usage) {
                html = "客户库位";
            } else if ("inventory" == usage) {
                html = "盘点库位";
            } else if ("procurement" == usage) {
                html = "补货库位";
            } else if ("production" == usage) {
                html = "生产库位";
            } else if ("transit" == usage) {
                html = "转移库位";
            } else {
                html = "状态未知";
            }
            return html;
        }
    },
    { title: "库位编码", field: 'Barcode', align: "center", sortable: true, order: "desc" },
    {
        title: "有效",
        field: 'Active',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.Active) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "退货库位",
        field: 'ReturnLocation',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.ReturnLocation) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },
    {
        title: "废料库位",
        field: 'ScrapLocation',
        align: "center",
        sortable: true,
        order: "desc",
        formatter: function cellStyle(value, row, index) {
            var html = "";
            if (row.ScrapLocation) {
                html = '<i class="fa fa-check"></i><span style="display:none;">是<span>';
            } else {
                html = '<i class="fa fa-remove"></i><span style="display:none;">否<span>';
            }
            return html;
        }
    },

    { title: "通道(X)", field: 'Posx', align: "center", sortable: true, order: "desc" },
    { title: "货架(Y)", field: 'Posy', align: "center", sortable: true, order: "desc" },
    { title: "层", field: 'Posz', align: "center", sortable: true, order: "desc" },

    {
        title: "操作",
        align: "center",
        field: 'action',
        formatter: function cellStyle(value, row, index) {

            var html = "";
            var url = "/stock/location/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
displayTable("#table-sale-order", "/sale/order", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "订单号", field: 'Name', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "创建时间", field: 'CreateDate', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "客户", field: 'Partner', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "业务员", field: 'SalesMan', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "所属公司", field: 'Company', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "发货仓库", field: 'StockWarehouse', align: "left", sortable: true, order: "desc", valign: "middle" },
    {
        title: "发货策略",
        field: 'PickingPolicy',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = "-";
            if (row.PickingPolicy == "one") {
                html = "一次发货";
            } else if (row.state == 'mult') {
                html = "分批发货";
            }
            return html;
        }
    },
    {
        title: "状态",
        field: 'State',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = "-";
            if (row.State == "draft") {
                html = "草稿";
            } else if (row.state == 'confirm') {
                html = "确认";
            } else if (row.state == 'cancel') {
                html = "取消";
            } else if (row.state == 'done') {
                html = "完成";
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
            var url = "/sale/order/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
displayTable("#table-purchase-order", "/purchase/order", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "订单号", field: 'Name', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "创建时间", field: 'CreateDate', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "供应商", field: 'Partner', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "采购员", field: 'PurchasesMan', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "所属公司", field: 'Company', align: "left", sortable: true, order: "desc", valign: "middle" },
    { title: "发货仓库", field: 'StockWarehouse', align: "left", sortable: true, order: "desc", valign: "middle" },
    {
        title: "状态",
        field: 'State',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = "-";
            if (row.State == "draft") {
                html = "草稿";
            } else if (row.state == 'confirm') {
                html = "确认";
            } else if (row.state == 'cancel') {
                html = "取消";
            } else if (row.state == 'done') {
                html = "完成";
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
            var url = "/purchase/order/";
            html += "<a href='" + url + row.id + "?action=edit' class='table-action btn btn-xs btn-default'>编辑<i class='fa fa-pencil'></i></a>";
            html += "<a href='" + url + row.id + "?action=detail' class='table-action btn btn-xs btn-default'>详情<i class='fa fa-external-link'></i></a>";
            return html;
        }
    }
]);
$.fn.editable.defaults.mode = 'inline';
$.fn.editable.defaults.emptytext = '无值';


//---------------------------------------款式中的属性列表----------------------------------
// 判断表单的状态
var formState = function(formIdSle) {
    'use strict';
    var form = $(formIdSle);
    if (form.hasClass("form-edit")) {
        return "edit"
    } else {
        return "readonly"
    }
};
//bootstrapTable
$("#one-product-template-attribute").bootstrapTable({
    method: "post",
    dataType: "json",
    locale: "zh-CN",
    contentType: "application/x-www-form-urlencoded",
    sidePagination: "server",
    url: "/product/template",
    dataField: "data",
    pagination: true,
    pageNumber: 1,
    pageSize: 20,
    pageList: [10, 25, 50, 100, 500, 1000],
    queryParams: function(params) {
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            params._xsrf = xsrf[0].value;
        }
        var recordId = $("input[name='_recordId']");
        if (recordId.length > 0) {
            params.recordId = recordId[0].value;
        }
        params.action = 'attribute';
        return params;
    },
    columns: [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        {
            title: "属性名称",
            field: 'name',
            sortable: true,
            order: "desc",
            formatter: function cellStyle(value, row, index) {
                var Attribute = row.Attribute;
                var html = "<p class='p-form-tree-disabled'>" + Attribute.name + "</p>";
                html += '<select name="productAttribute" id="productAttributeID-' + index + '" data-action="' + row.action + '"class="form-control select-product-attribute">' +
                    '</select>';
                return html;
            }
        },
        {
            title: "属性值",
            field: 'AttributeValues',
            sortable: true,
            formatter: function cellStyle(value, row, index) {
                var attributeValues = row.AttributeValues;
                var html = "";
                html += "<p class='p-form-tree-disabled'>";
                for (line in attributeValues) {
                    html += "<a class='display-block label label-primary'>" + line.name + "</a>";
                }
                html += "</p>";
                html += '<select name="productAttributeValue" id="productAttributeValueID-' + index + '"data-action="' + row.action + '" class="form-control select-product-attribute-value" multiple="multiple" >'
                for (line in attributeValues) {
                    html += "<a class='display-block label label-primary'>" + line.name + "</a>";
                }
                for (line in attributeValues) {
                    html += '<option value="' + line.id + '" selected="selected">' + line.name + '</option>';
                }
                html += '</select>';
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
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-danger'>删除&nbsp<i class='fa  fa-trash'></i></a>";
                return html;
            }
        }
    ],
});
//x-editable
$(".form-table-add-line").on("click", function(e) {
    var formId = e.currentTarget.dataset["formid"];
    $("#one-product-template-attribute").bootstrapTable("append", randomData());
    select2AjaxData(".select-product-attribute", '/product/attribute/?action=search'); // 选择属性
    select2AjaxData(".select-product-attribute-value", '/product/attributevalue/?action=search'); // 选择属性值
    function randomData(e) {
        rows = [];
        rows.push({
            action: "create",
            id: 0,
            Attribute: "",
            AttributeValues: []
        });
        return rows;
    }
});
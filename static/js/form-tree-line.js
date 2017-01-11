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
                var attribute = row.Attribute;
                var html = "<p class='p-form-tree-disabled'>" + attribute.name + "</p>";
                html += '<select name="productAttributeID" id="productAttributeID" class="form-control select-product-attribute">' +
                    '</select>';

                return html;
            }
        },
        {
            title: "属性值",
            field: 'attributes',
            sortable: true,
            formatter: function cellStyle(value, row, index) {
                var html = "";

                var attributeValues = row.AttributeValues;
                for (line in attributeValues) {

                }
                return html;
            }
        },
    ],
});
//x-editable
$(".form-table-add-line").on("click", function(e) {
    var formId = e.currentTarget.dataset["formid"];
});
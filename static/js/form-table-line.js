displayTable("#form-table-sale-order-line", "/sale/order/line", [
    { title: "全选", field: 'ID', checkbox: true, align: "center", valign: "middle" },
    { title: "订单明细号", field: 'Name', align: "left", sortable: true, order: "desc", valign: "middle" },
    {
        title: "产品",
        field: 'Product',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.Product) {
                html = "<p class='p-form-line-control'>" + "[" + row.Product.defaultCode + "]" + row.Product.name + "<a class='pull-right' href='/product/product/" + row.Product.id + "?action=detail'><i class='fa fa-external-link'></i></a></p>";
                html += '<select data-type="int" data-oldValue="' + row.Product.id + '" name="ProductProduct-' + row.id + '" id="ProductProduct-' + row.id + '" class="form-control select-sale-order-product-product">';
                html += '<option value="' + row.Product.id + '"  selected="selected">' + '[' + row.Product.defaultCode + ']' + row.Product.name + '</option>'
                html += '</select>';
            }
            return html;
        }
    },
    {
        title: "产品编码",
        field: 'ProductCode',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.ProductCode) {
                html = "<p class='p-form-line-control'>" + row.ProductCode + "</p>";
            }
            return html;
        }
    },
    {
        title: "产品名称",
        field: 'ProductName',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.ProductName) {
                html = "<p class='p-form-line-control'>" + row.ProductName + "</p>";
            }
            return html;
        }
    },
    {
        title: "第一单位数量",
        field: 'FirstSaleQty',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.FirstSaleQty) {
                html = "<p class='p-form-line-control'>" + row.FirstSaleQty + "</p>";
            }
            return html;
        }
    },
    {
        title: "第二单位数量",
        field: 'SecondSaleQty',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.SecondSaleQty) {
                html = "<p class='p-form-line-control'>" + row.SecondSaleQty + "</p>";
            }
            return html;
        }
    },
    {
        title: "单价",
        field: 'PriceUnit',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.PriceUnit) {
                html = "<p class='p-form-line-control'>" + row.PriceUnit + "</p>";
            }
            return html;
        }
    },
    {
        title: "小计",
        field: 'Total',
        align: "left",
        sortable: true,
        order: "desc",
        valign: "middle",
        formatter: function cellStyle(value, row, index) {
            var html = '';
            if (row.Total) {
                html = "<p class='p-form-line-control'>" + row.Total + "</p>";
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
            var url = "/sale/order/line/";
            html += "<i class='fa fa-trash-o'></i>";
            return html;
        }
    }
], undefined, function() {
    select2AjaxData(".select-sale-order-product-product", '/product/product/', function(event) {});
});
// 增加一行销售订单明细
$("#add-one-sale-order-line").on("click", function(e) {
    $("#form-table-sale-order-line").bootstrapTable('prepend', [{
        FirstSaleQty: 1,
        ID: null,
        Name: "",
        PriceUnit: 0,
        Product: {
            id: null,
            name: "",
            defaultCode: ""
        },
        ProductCode: "",
        ProductName: "",
        SecondSaleQty: 0,
        Total: 0,
        id: null
    }]);
});
$(function() {
    var saleOrderState = $(".sale-order-state");
    if (saleOrderState.length > 0) {
        var params = {};
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            params._xsrf = xsrf[0].value;
        }
        var recordID = $("input[name ='recordID']");
        if (recordID.length > 0) {
            params.recordID = recordID[0].value;
        }
        var companyID = $("input[name ='Company']");
        if (companyID.length > 0) {
            params.companyID = companyID[0].value;
        }
        var stockWarehouseID = $("input[name ='StockWarehouse']");
        if (stockWarehouseID.length > 0) {
            params.StockWarehouseID = stockWarehouseID[0].value;
        }
        $.ajax({
            type: "POST",
            dataType: "json",
            url: "/sale/order/state/?action=search",
            data: params,
            success: function(result) {},
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                console.log(XMLHttpRequest.status);
                console.log(XMLHttpRequest.readyState);
                console.log(textStatus);
                toastr.warning("订单状态获取失败", "警告");
            }
        });
    }
});
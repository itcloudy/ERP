$(function() {
    'use strict';
    var displayStockPickingTypeKanban = function(el, dataArr) {
        var innerHtml = "";
        for (var i = 0, len = dataArr.length; i < len; i++) {
            var type = "";
            if (dataArr[i].Code == "outgoing") {
                type = "出库";

            } else if (dataArr[i].Code == "incoming") {
                type = "入库";
            } else if (dataArr[i].Code == "internal") {
                type = "内部调拨";
            }
            innerHtml += '<div class="col-md-4">';
            innerHtml += '<div class="box box-success">';
            innerHtml += '<div class="box-header with-border">';
            innerHtml += '<h3 class="box-title pull-left"><a class="text-primary" href="/stock/warehouse/' + dataArr[i].WareHouse.id + '?action=detail">' + dataArr[i].WareHouse.name + '[' + type + ']' + ':</a></h3>';
            innerHtml += '<h3 class="box-title">&nbsp&nbsp&nbsp<a class="text-danger" href="/stock/picking/type/' + dataArr[i].id + '?action=detail">' + dataArr[i].Name + '</a></h3>';
            innerHtml += '<div class="box-tools pull-right">';
            innerHtml += '<button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i></button>';
            innerHtml += '<button type="button" class="btn btn-box-tool" data-widget="remove"><i class="fa fa-times"></i></button>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '<div class="box-body">';
            innerHtml += '<div class="row">';
            innerHtml += '<div class="col-md-12">';
            innerHtml += "等待添加内容";
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
        }
        el.append(innerHtml);
    };
    var stockPickingTypeKanban = $("#kanban-stock-picking-type");
    if (stockPickingTypeKanban.length > 0) {
        $.ajax({
            type: "POST",
            url: "/stock/picking/type/",
            dataType: "json",
            data: (function() {
                var params = {
                    action: "search",
                };
                var xsrf = $("input[name ='_xsrf']");
                if (xsrf.length > 0) {
                    xsrf = xsrf[0].value;
                    params._xsrf = xsrf;
                }
                return params;
            })(),
            success: function(response) {
                var data = response.data;
                if (data != undefined && data.length > 0) {
                    displayStockPickingTypeKanban(stockPickingTypeKanban, data);
                }
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                console.log(XMLHttpRequest.status);
                console.log(XMLHttpRequest.readyState);
                console.log(textStatus);
                toastr.error("请求失败，请刷新页面", "错误");
            }
        });
    }
    var displaySaleCounterKanban = function(el, dataArr) {
        var innerHtml = "";
        for (var i = 0, len = dataArr.length; i < len; i++) {

            innerHtml += '<div class="col-md-4">';
            innerHtml += '<div class="box box-success">';
            innerHtml += '<div class="box-header with-border">';
            innerHtml += '<h3 class="box-title pull-left"><a class="text-primary" href="/company/' + dataArr[i].Company.id + '?action=detail">' + dataArr[i].Company.name + ':</a></h3>';
            innerHtml += '<h3 class="box-title">&nbsp&nbsp&nbsp<a class="text-danger" href="/sale/counter/' + dataArr[i].id + '?action=detail">' + dataArr[i].Name + '</a></h3>';
            innerHtml += '<div class="box-tools pull-right">';
            innerHtml += '<button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i></button>';
            innerHtml += '<button type="button" class="btn btn-box-tool" data-widget="remove"><i class="fa fa-times"></i></button>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '<div class="box-body">';
            innerHtml += '<div class="row">';
            innerHtml += '<div class="col-md-12">';
            innerHtml += "等待添加内容";
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
            innerHtml += '</div>';
        }
        el.append(innerHtml);
    };
    var saleCounterKanban = $("#kanban-sale-counter");
    if (saleCounterKanban.length > 0) {
        $.ajax({
            type: "POST",
            url: "/sale/counter/",
            dataType: "json",
            data: (function() {
                var params = {
                    action: "search",
                };
                var xsrf = $("input[name ='_xsrf']");
                if (xsrf.length > 0) {
                    xsrf = xsrf[0].value;
                    params._xsrf = xsrf;
                }
                return params;
            })(),
            success: function(response) {
                var data = response.data;
                if (data != undefined && data.length > 0) {
                    displaySaleCounterKanban(saleCounterKanban, data);
                }
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                console.log(XMLHttpRequest.status);
                console.log(XMLHttpRequest.readyState);
                console.log(textStatus);
                toastr.error("请求失败，请刷新页面", "错误");
            }
        });
    }

});
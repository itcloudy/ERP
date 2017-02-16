$(function() {
    'use strict';
    var displayKanban = function(el, dataArr) {
        /*<div class="col-md-12">
            <div class="box box-success">
                <div class="box-header with-border box-customer">
                    <h3 class="box-title text-success">客户</h3>
                    <div class="box-tools pull-right">
                        <button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i></button>
                        <button type="button" class="btn btn-box-tool" data-widget="remove"><i class="fa fa-times"></i></button>
                    </div>
                </div>
                <div class="box-body">
                    <div class="row">
                        <div class="col-md-12">
                            <p>body</p>
                        </div>
                    </div>
                </div>
                <div class="box-footer">
                    <div class="row">
                        <div class="col-md-12">
                            <p>foolter</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>*/
        var innerHtml = "";
        for (var i = 0, len = dataArr.length; i < len; i++) {
            innerHtml += '<div class="col-md-3">';
            innerHtml += '<div class="box box-success">';
            innerHtml += '<div class="box-header with-border">';
            innerHtml += '<h3 class="box-title text-primary pull-left">' + dataArr[i].WareHouse.name + ':</h3>';
            innerHtml += '<h3 class="box-title text-danger ">&nbsp&nbsp&nbsp' + dataArr[i].Name + '</h3>';
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
    var stockPickingType = $("#kanban-stock-picking-type");
    if (stockPickingType.length > 0) {
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
                    displayKanban(stockPickingType, data);
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
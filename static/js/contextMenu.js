$(function() {
    // 'table-product-product'
    //对产品规格数据进行处理
    $.contextMenu({
        selector: '#table-product-product tr.selected',
        build: function($trigger, e) {
            return {
                callback: function(key, options) {
                    var m = "clicked: " + key;
                    console.log(m);
                    console.log(options);
                    console.log($("#table-product-product").bootstrapTable('getSelections'))
                    var selectedArr = $("#table-product-product").bootstrapTable('getSelections');
                    var selectedIds = [];
                    var len = selectedArr.length;
                    if (len > 0) {
                        for (var i = 0; i < len; i++) {
                            selectedIds.push(selectedArr[i].id);
                        }
                        var params = {
                            action: 'batchUpdate',
                            ids: selectedIds
                        };
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf != undefined) {
                            params._xsrf = xsrf[0].value;
                        }
                        if (key == "activeFalse") {
                            params.field = "Active";
                            params.value = false;

                        } else if (key == "activeTrue") {
                            params.field = "Active";
                            params.value = true;
                        }
                        $.ajax({
                            type: "POST",
                            url: "/product/product/",
                            data: params,
                            dataType: "json",
                            success: function(response) {
                                if (response.code == 'failed') {
                                    toastr.error("修改失败", "错误");
                                    return;
                                } else {
                                    toastr.success("请刷新页面更新数据", "修改成功");
                                    return;
                                }
                            },
                            error: function(XMLHttpRequest, textStatus, errorThrown) {
                                console.log(XMLHttpRequest.status);
                                console.log(XMLHttpRequest.readyState);
                                console.log(textStatus);
                                toastr.error("请求失败，请刷新页面后再操作", "错误");
                            }
                        });
                    }
                },
                items: {
                    "normalSub": {
                        icon: "edit",
                        name: "修改",
                        items: {
                            "activeFalse": { name: "下架" },
                            "activeTrue": { name: "上架" },
                        }
                    }
                }
            };
        }
    });
});
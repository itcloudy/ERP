$(function() {
    // 'table-product-product'
    //对产品规格数据进行处理
    $.contextMenu({
        selector: '#table-product-product tr.selected',
        build: function($trigger, e) {
            return {
                callback: function(key, options) {
                    var $selector = $("#table-product-product");
                    if ("excel" != key) {
                        var selectedArr = $selector.bootstrapTable('getSelections');
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
                                        if (key == "activeTrue") {
                                            $("#table-product-product tr.selected td.data-active .fa-remove").removeClass("fa-remove").addClass("fa-check");
                                        } else {
                                            $("#table-product-product tr.selected td.data-active .fa-check").removeClass("fa-check").addClass("fa-remove");
                                        }
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
                    } else {
                        $selector.tableExport({
                            type: key,
                            escape: false
                        });
                    }
                },
                items: {
                    "edit": {
                        icon: "edit",
                        name: "修改",
                        items: {
                            "activeTrue": { icon: "fa-arrow-up", name: "上架" },
                            "activeFalse": { icon: "fa-arrow-down", name: "下架" }
                        }
                    },
                    "export": {
                        icon: "fa-download",
                        name: "导出",
                        items: {
                            "excel": { icon: "fa-file-excel-o", name: "Excel" },
                        }
                    }
                }
            };
        }
    });
});
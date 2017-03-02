$(function() {
    'use strict';
    // 左侧菜单显示
    var menuDisplay = function() {
        var currentMenu = $("#top-menu li.active");
        if (currentMenu.length > 0) {
            currentMenu = $(currentMenu)[0];
            var parentNode = currentMenu.parentNode;
            if (parentNode) {
                while (true) {
                    parentNode = $(parentNode);
                    if (parentNode.attr("id") == "top-menu") {
                        break;
                    } else {
                        parentNode = parentNode[0].parentNode;
                        if (parentNode) {
                            if (parentNode.nodeName == "LI") {
                                $(parentNode).addClass("active");
                            }
                        }
                    }
                }
            }
        }
    };
    menuDisplay();
    // 单选checkbox
    $('.form-checkbox').each(function(index, el) {
        $(el).iCheck({
            checkboxClass: 'icheckbox_square-green',
            radioClass: 'iradio_square-green',
            increaseArea: '20%',
        });
        if ($(el).attr('checked') == 'checked') {
            $(el).iCheck("check");
        } else {
            if ($(el).data("oldvalue") === "false" || $(el).data("oldvalue") === "" || $(el).data("oldvalue") == undefined) {
                $(el).iCheck("uncheck");
            } else {
                $(el).iCheck("check");
            }
        }

        var form = $(el)[0].form;
        if (form != undefined) {
            if ($(form).hasClass("form-disabled")) {
                $(el).iCheck("disable");
            }
        }
    });

    $('.input-radio').each(function(index, el) {
        $(el).iCheck({
            checkboxClass: 'icheckbox_square-green',
            radioClass: 'iradio_square-green',
            increaseArea: '20%',
        });
        if ($(el).data("oldvalue") == $(el).val()) {
            $(el).iCheck("check");
        } else {
            $(el).iCheck("uncheck");
        }
        var form = $(el)[0].form;
        if (form != undefined) {
            if ($(form).hasClass("form-disabled")) {
                $(el).iCheck("disable");
            }
        }
    });

    // //编辑删除readonly属性，输入框变成可编辑状态
    // $(".form-edit-btn").on("click", function(e) {
    //     e.preventDefault();
    //     $(".input-radio").iCheck("enable");
    //     $(".form-disabled").addClass("form-edit").removeClass("form-disabled");

    // });
    // $(".form-cancel-btn").on("click", function(e) {
    //     e.preventDefault();
    //     $(".input-radio").iCheck("disable");
    //     $(".form-edit").addClass("form-disabled").removeClass("form-edit");
    // });
    $(".select-product-uom-category-type").on("change", function(e) {
        if (e.currentTarget.value == "1") {
            $("#factorInvDisplay").addClass("hidden");
            $("#factorDisplay").removeClass("hidden");
        } else if (e.currentTarget.value == "3") {
            $("#factorDisplay").addClass("hidden");
            $("#factorInvDisplay").removeClass("hidden");
        } else {
            $("#factorDisplay").addClass("hidden");
            $("#factorInvDisplay").addClass("hidden");
        }
    });

    //如果搜索添加不为空，增加提示样式
    $("#listViewSearch .filter-condition").change(function(e) {
        e.currentTarget.value = e.currentTarget.value.trim();
        var nums = $.grep($("#listViewSearch input"), function(el, index) {
            if (el.value != "") {
                return true;
            } else {
                return false;
            }
        });
        if (nums.length > 0) {
            if ($("button[id^='clearListSearchCond']:first").hasClass("hide")) {
                $("button[id^='clearListSearchCond']").removeClass("hide");
            }
        } else {
            if (!$("button[id^='clearListSearchCond']:first").hasClass("hide")) {
                $("button[id^='clearListSearchCond']").addClass("hide");
            }
        }
    });
    // 若过滤条件不为空， 显示清空条件按钮
    (function() {
        var nums = $.grep($("#listViewSearch input"), function(el, index) {
            if (el.value != "") {
                return true;
            } else {
                return false;
            }
        });
        // console.log(nums);
        if (nums.length < 1) {
            $("button[id^='clearListSearchCond']").addClass("hide");
        } else {
            $("button[id^='clearListSearchCond']").removeClass("hide");
        }
    })();
    $("button[id^='clearListSearchCond']").click(function(e) {
        $("#listViewSearch input").each(function() {
            this.value = "";
        });
        $(this).addClass("hide");
        $(".table-diplay-info").bootstrapTable('refresh');
    });

});
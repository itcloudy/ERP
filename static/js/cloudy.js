$(document).ready(function() {
    'use strict';
    $('.input-radio').iCheck({
        checkboxClass: 'icheckbox_square-green',
        radioClass: 'iradio_square-green',
        increaseArea: '20%'
    });
    //有checked的radio默认选中
    $("input.checked").iCheck("check");
    //编辑删除readonly属性，输入框变成可编辑状态
    $(".form-edit-btn").on("click", function(e) {
        e.preventDefault();
        $(".input-radio").iCheck("enable");
        $(".form-disabled").addClass("form-edit").removeClass("form-disabled");
    });
    $(".form-cancel-btn").on("click", function(e) {
        e.preventDefault();
        $(".input-radio").iCheck("disable");
        $(".form-edit").addClass("form-disabled").removeClass("form-edit");
    });
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
    $("#listViewSearch input").change(function(e) {
        e.currentTarget.value = e.currentTarget.value.trim();
        var nums = $.grep($("#listViewSearch input"), function(el, index) {
            if (el.value != "") {
                return true
            } else {
                return false
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
                return true
            } else {
                return false
            }
        });
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
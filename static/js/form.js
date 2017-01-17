$(function() {

    // 保存事件处理
    $(".form-save-btn").on("click", function(e) {

        var formData = {};
        var form = e.currentTarget.form;
        // form表单验证
        $('#' + form.id).bootstrapValidator('validate');
        //    获得form直接的字段
        var formFields = $(form).find(".form-create,.form-edit");
        console.log(formFields);
        for (var i = 0, len = formFields.length; i < len; i++) {
            // 处理radio数据
            var self = formFields[i];
            if (self.type == "radio") {
                if ($(self).hasClass("checked")) {
                    formData[self.name] = $(self).val();
                }
            } else if (self.type == "checkbox") {
                if (self.checked) {
                    formData[self.name] = true;
                }
            } else {
                var val = $(self).val();
                if (val != "") {
                    formData[self.name] = val;
                }
            }
        }
        // console.log(formData);
        e.preventDefault();

    });
    //文件导入
    $('#import-file-excel').fileinput({
        language: 'zh',
        uploadUrl: '#',
        multiple: false,
        minFileCount: 1,
        showPreview: false,
        uploadExtraData: (function() {
            'use strict';
            var params = {};
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf.length > 0) {
                params._xsrf = xsrf[0].value;
            }
            params.upload = "uploadFile";
            params.action = "upload";
            params._method = "PUT";
            return params;
        })(),
        allowedFileExtensions: ['xlsx', 'csv', 'xls'],
    });
    // 图片上传处理
    $('#product-images').fileinput({
        language: 'zh',
        uploadUrl: '#',
        uploadExtraData: (function() {
            var params = {};
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf.length > 0) {
                params._xsrf = xsrf[0].value;
            }
            params.upload = "uploadFile";
            params.action = "upload";
            params._method = "PUT";
            return params;
        })(),
        allowedFileExtensions: ['jpg', 'png', 'gif'],
    });
    $(".form-disabled .file-input").hide();
    $("#productTemplateForm .form-edit-btn").bind("click.images", function() {
        $(".file-input").show();
    });
    $("#productTemplateForm .form-save-btn,#productTemplateForm .form-cancel-btn").bind("click.images", function() {
        $(".file-input").hide();
    });
    // 单击图片悬浮
    $(".click-modal-view").dblclick(function(e) {
        var images = $(".click-modal-view");
        var imagesLen = images.length;
        var indicatorsHtml = "";
        var carouselInnerHtml = "";
        for (var i = 0; i < imagesLen; i++) {
            if (i == 0) {
                indicatorsHtml += ' <li data-target="#productImagesCarousel" data-slide-to=' + i + ' class="active"></li>';
                carouselInnerHtml += '<div class="item active"> <img src="' + images[i].src + '" alt=""> </div>';
            } else {
                indicatorsHtml += ' <li data-target="#productImagesCarousel" data-slide-to=' + i + '></li>';
                carouselInnerHtml += '<div class="item "> <img src="' + images[i].src + '" alt=""> </div>';
            }
        }
        console.log(indicatorsHtml);
        $("#productImagesCarousel .carousel-indicators").append(indicatorsHtml);
        $("#productImagesCarousel .carousel-inner").append(carouselInnerHtml);
        $('#productImagesModal').modal('show');
    });
    // 款式form中图片懒加载
    $('a[href="#productImages"]').on('shown.bs.tab', function(e) {
        // 图片加载
        $("#productImages .click-modal-view").each(function(index, el) {
            if ($(el).attr("src") == "") {
                $(el).attr("src", $(el)[0].dataset["src"]);
            }
        });
    });
});
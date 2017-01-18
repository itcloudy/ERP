$(function() {
    // 保存事件处理
    $(".form-save-btn").on("click", function(e) {
        var form = e.currentTarget.form;
        var formNode = $('#' + form.id);
        var formData = {
            FormAction: "create"
        };
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            formData._xsrf = xsrf[0].value;
        }
        // form表单验证
        var bootstrapValidator = formNode.data('bootstrapValidator');
        //手动触发验证
        bootstrapValidator.validate();
        // 验证结果
        var formValid = bootstrapValidator.isValid();
        if (!formValid) {
            toastr.error("数据验证失败，请检查数据", "错误");
            return;
        }
        //    获得form直接的字段
        var formFields = $(form).find(".form-create,.form-edit");
        if ($(form).find("input[name='recordID']").length > 0) {
            form.FormAction = "update";
        }
        for (var i = 0, len = formFields.length; i < len; i++) {
            var self = formFields[i];
            // 处理radio数据
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
        //获得form-tree-create信息
        var formTreeFields = $(form).find(".form-tree-line-create");
        for (var i = 0, lineLen = formTreeFields.length; i < lineLen; i++) {
            console.log(lineLen);
            var self = formTreeFields[i];
            var treeName = $(self).data("treename");
            console.log(formData[treeName]);
            if (formData[treeName] == undefined); {
                formData[treeName] = [];
            }
            var treeArray = [];
            var cellFields = $(self).find(".form-tree-create");
            var cellData = {
                FormAction: "create"
            };
            var hasProp = false;
            for (var j = 0, cellLen = cellFields.length; j < cellLen; j++) {
                var cell = cellFields[j];
                var cellName = cell.name;
                var cellValue = $(cell).val();
                if (cellValue != null) {
                    cellData[cellName] = cellValue;
                    hasProp = true;
                }
            }
            if (hasProp) {
                console.log(cellData);
                treeArray = treeArray.push(cellData);
            }
        }
        console.log(formData);
        $.post(form.action, formData).success(function(response) {
            if (response.code == 'failed') {
                toastr.error("创建失败", "错误");
                return;
            } else {
                toastr.success("创建成功");
                // window.location = response.location;
            }
        });
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
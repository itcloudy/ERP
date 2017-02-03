var BootstrapValidator = function(selector, needValidatorFields) {

    $(selector).bootstrapValidator({
        message: '该值无效',
        feedbackIcons: { /*input状态样式图片*/
            valid: 'glyphicon glyphicon-ok',
            invalid: 'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        live: 'enabled',
        submitButtons: 'button[type="submit"]',
        trigger: null,
        fields: needValidatorFields
    }).on('success.form.bv', function(e) {
        // Prevent form submission
        e.preventDefault();


        var $form = $(e.currentTarget);
        var formData = {
            FormAction: "create" //默认为创建
        };
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            xsrf = xsrf[0].value;
            // formData._xsrf = xsrf[0].value;
        }
        //    获得form直接的字段
        var formFields = $form.find(".form-create,.form-edit");
        if ($form.find("input[name='recordID']").length > 0) {
            formData.FormAction = "update";
        }
        //根据数据类型获得正确的数据,默认string
        var getCurrentDataType = function(val, dataType) {
            if (dataType == "" || dataType === undefined || dataType === null) {
                dataType = "string";
            }
            switch (dataType) {
                case "int": // 整形
                    val = parseInt(val);
                    break;
                case "float": // 浮点型
                    val = parseFloat(val);
                    break;
                case "array_int": // 整形数组
                    var a_arr = [];
                    for (var a_i = 0, a_l = val.length; a_i < a_l; a_i++) {
                        a_arr.push(parseInt(val[a_i]));
                    }
                    val = a_arr;
                    break;
                case "arrar_float": //  浮点型数组
                    var a_arr = [];
                    for (var a_i = 0, a_l = val.length; a_i < a_l; a_i++) {
                        a_arr.push(parseFloat(val[a_i]));
                    }
                    val = a_arr;
                    break;
            }
            return val
        };
        for (var i = 0, len = formFields.length; i < len; i++) {
            var self = formFields[i];
            var oldValue = null;
            // console.log(self.name + ":" + $(self).val());
            oldValue = $(self).data("oldvalue");
            // 处理radio数据
            if (self.type == "radio") {
                if ($(self).data("type") == "string") {
                    if ($(self).hasClass("checked")) {
                        formData[self.name] = $(self).val();
                    }
                }
            } else if (self.type == "checkbox") {
                if (self.checked) {
                    formData[self.name] = true;
                } else {
                    formData[self.name] = false;
                }
            } else {
                var val = $(self).val();

                if (val != "") {
                    // 若为null跳出此次循环
                    if (val === null) {
                        continue;
                    }
                    formData[self.name] = getCurrentDataType(val, $(self).data("type"))
                }
            }
        }
        var getTreeLineData = function(cellFields, action = "create") {
            var funCellData = {
                FormAction: action
            };
            for (var j = 0, cellLen = cellFields.length; j < cellLen; j++) {
                var funHasProp = false;
                var cell = cellFields[j];
                var cellName = cell.name;
                var oldValue = $(cell).data("oldvalue");
                var cellValue = $(cell).val();
                if (cellValue != "") {
                    if (cellValue === null) {
                        continue;
                    }
                    funCellData[cellName] = getCurrentDataType(cellValue, $(cell).data("type"));
                    funHasProp = true;
                }
            }
            return { cellData: funCellData, hasProp: funHasProp };
        };
        //获得form-tree-create信息
        var formTreeFields = $form.find(".form-tree-line-create");
        for (var i = 0, lineLen = formTreeFields.length; i < lineLen; i++) {
            var self = formTreeFields[i];
            var treeName = $(self).data("treename");
            if (formData[treeName] == undefined) {
                formData[treeName] = [];
            }
            var cellFields = $(self).find(".form-line-cell-create");
            var resultCreate = getTreeLineData(cellFields, "create");
            if (resultCreate.hasProp) {
                formData[treeName].push(resultCreate.cellData);
            }
        }
        //获得form-tree-edit信息
        var formTreeFields = $form.find(".form-tree-line-edit");
        for (var i = 0, lineLen = formTreeFields.length; i < lineLen; i++) {
            var self = formTreeFields[i];
            var treeName = $(self).data("treename");
            if (formData[treeName] == undefined) {
                formData[treeName] = [];
            }
            var createCellFields = $(self).find(".form-line-cell-create");
            if (createCellFields.length > 0) {
                var resultCreate = getTreeLineData(createCellFields, "create");
                if (resultCreate.hasProp) {
                    formData[treeName].push(resultCreate.cellData);
                }
            }
            var editCellFields = $(self).find(".form-line-cell-edit");
            if (editCellFields.length > 0) {
                var resultCreate = getTreeLineData(editCellFields, "update");
                if (resultCreate.hasProp) {
                    formData[treeName].push(resultCreate.cellData);
                }
            }
        }
        var postParams = {
            postData: JSON.stringify(formData),
            _xsrf: xsrf
        };
        var method = $form.find("input[name='_method']");
        if (method.length > 0) {
            postParams._method = method.val();
        }
        console.log(postParams);
        $.post($form.action, postParams).success(function(response) {
            if (response.code == 'failed') {
                if (formData.FormAction == "update") {
                    toastr.error("修改失败", "错误");
                } else {
                    toastr.error("创建失败", "错误");
                }
                return;
            } else {
                if (formData.FormAction == "update") {
                    toastr.success("<h3>修改成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
                } else {
                    toastr.success("<h3>创建成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
                }
                // setTimeout(function() { window.location = response.location; }, 1000);
            }
        });
        // Use Ajax to submit form data
        // $.post($form.attr('action'), $form.serialize(), function(result) {
        //     console.log(result);
        // }, 'json');
    });
};

$(function() {
    'use strict';
    // 用户form
    BootstrapValidator("#userForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "用户名不能为空"
                },
                stringLength: {
                    min: 3,
                    max: 20,
                    message: '用户名长度必须在3到20之间'
                },
                remote: {
                    url: "/user/",
                    message: "用户已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {

                        var params = {
                            action: "validator"
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var name = $('input[name="name"]');
                        if (name.length > 0) {
                            params.name = name[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
                regexp: {
                    regexp: /^[a-zA-Z0-9_\.]+$/,
                    message: '用户名由数字字母下划线和.组成'
                }
            },
        },
        NameZh: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "用户名(中文)不能为空"
                }
            },
        },
        Mobile: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "手机号码不能为空"
                },
                remote: {
                    url: "/user/",
                    message: "手机号码已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {

                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var mobile = $('input[name="mobile"]');
                        if (mobile.length > 0) {
                            //用户名，手机号码，邮箱都必须唯一
                            params.name = mobile[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        },
        Email: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "邮箱不能为空"
                },
                remote: {
                    url: "/user/",
                    message: "邮箱已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var email = $('input[name="email"]');
                        if (email.length > 0) {
                            //用户名，手机号码，邮箱都必须唯一
                            params.name = email[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
                regexp: {
                    regexp: /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/,
                    message: '邮箱地址无效'
                }
            }
        },
        Position: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "职位不能为空"
                }
            }
        },
        Department: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "部门不能为空"
                }
            }
        },
        Group: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "权限组不能为空"
                }
            }
        },
        Password: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "密码不能为空"
                }
            }
        },
    });
    // 序列号
    BootstrapValidator("#sequenceForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: { message: "名称不能为空" },
                remote: {
                    url: "/sequence/",
                    message: "该序列名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    }
                }
            }
        },
        Prefix: {
            message: "该值无效",
            validators: {
                notEmpty: { message: "前缀不能为空" },
                remote: {
                    url: "/sequence/",
                    message: "该前缀已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    }
                },
                regexp: { /* 只需加此键值对，包含正则表达式，和提示 */
                    regexp: /[A-Z]+$/,
                    message: '只能是A-Z的大写字母'
                }
            }
        },
        StructName: {
            message: "该值无效",
            validators: {
                notEmpty: { message: "表struct名称不能为空" },
                remote: {
                    url: "/sequence/",
                    message: "该表struct名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    }
                }
            }
        }
    });
    //产品分类form
    BootstrapValidator("#productCategoryForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "产品类别不能为空"
                },
                remote: {
                    url: "/product/category/",
                    message: "该类别已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            },
        },
    });
    //产品属性form
    BootstrapValidator("#productAttributeForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "属性名称不能为空"
                },
                remote: {
                    url: "/product/attribute/",
                    message: "该属性名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params;
                    },
                },
            },
        }
    });
    //产品属性值form
    BootstrapValidator("#productAttributeValueForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "属性值不能为空"
                },
                remote: {
                    url: "/product/attributevalue/",
                    message: "该属性值已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var attributeId = $("select[name='productAttributeID']");

                        if (attributeId.length > 0) {
                            attributeId = attributeId[0].value;
                            params.attributeId = attributeId;
                        }
                        return params
                    },
                },
            },
        },
        AttributeID: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "属性不能为空"
                }
            }
        }
    });
    //计量单位分类
    BootstrapValidator("#productUomCategForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "计量单位分类不能为空"
                },
                remote: {
                    url: "/product/uomcateg/",
                    message: "该计量单位分类已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        }
    });
    //计量单位
    BootstrapValidator("#productUomForm", {
        Category: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "计量单位分类不能为空"
                },
            }
        },
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "计量单位名称不能为空"
                },
                remote: {
                    url: "/product/uom/",
                    message: "该计量单位名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']")
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        },
        Rounding: {
            message: "该值无效",
            validators: {
                numeric: {
                    message: "舍入精度应为数字"
                },
            }
        },
        FactorInv: {
            message: "该值无效",
            validators: {
                numeric: {
                    message: "更大比率应为数字"
                },
            }
        },
        Factor: {
            message: "该值无效",
            validators: {
                numeric: {
                    message: "比率应为数字"
                },
            }
        }

    });

    //产品款式
    BootstrapValidator("#productTemplateForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "款式名称不能为空"
                },
                remote: {
                    url: "/product/template/",
                    message: "该款式名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var xsrf = $("input[name ='_xsrf']");
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        },
        StandardPrice: {
            validators: {
                numeric: {
                    message: "成本价格必须为数字"
                }
            }
        },
        Category: {
            validators: {
                notEmpty: {
                    message: "款式类别不能为空"
                }
            }
        },
        FirstSaleUom: {
            validators: {
                notEmpty: {
                    message: "销售单位1不能为空"
                }
            }
        },
        FirstPurchaseUom: {
            validators: {
                notEmpty: {
                    message: "采购单位1不能为空"
                }
            }
        },
        AttributeLineId: {
            validators: {
                notEmpty: {
                    message: "款式属性不能为空"
                }
            }
        },
        AttributeLineValueId: {
            validators: {
                notEmpty: {
                    message: "款式属性值不能为空"
                }
            }
        }
    });
    // 产品规格
    BootstrapValidator("#productProductForm", {
        ProductTemplateID: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "产品款式名称不能为空"
                },
            }
        },
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "产品规格名称不能为空"
                }
                // remote: {
                //     url: "/product/product/",
                //     message: "该产品规格名称已经存在",
                //     dataType: "json",
                //     delay: 200,
                //     type: "POST",
                //     data: function() {
                //         var params = {
                //             action: "validator",
                //         }
                //         var xsrf = $("input[name ='_xsrf']")
                //         if (xsrf.length > 0) {
                //             params._xsrf = xsrf[0].value;
                //         }
                //         var recordID = $("input[name ='recordID']");
                //         if (recordID.length > 0) {
                //             params.recordID = recordID[0].value;
                //         }
                //         return params
                //     },
                // }
            },
        },
        AttributeValueIds: {
            message: "该值无效",
            validators: {
                remote: {
                    url: "/product/product/",
                    message: "该产品款式属性值组合的产品规格已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var ProductTemplateID = $("#ProductTemplateID").val();
                        if (ProductTemplateID != undefined) {
                            params.ProductTemplateID = parseInt(ProductTemplateID);
                        }
                        var xsrf = $("input[name ='_xsrf']")
                        if (xsrf.length > 0) {
                            params._xsrf = xsrf[0].value;
                        }
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }

        },
        FirstSaleUom: {
            validators: {
                notEmpty: {
                    message: "销售单位1不能为空"
                }
            }
        },
        FirstPurchaseUom: {
            validators: {
                notEmpty: {
                    message: "采购单位1不能为空"
                }
            }
        }
    });
});
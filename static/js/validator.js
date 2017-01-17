// $(function() {
// 'use strict';
//用户
$("#userForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
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
                        var recordID = $("input[name ='_recordID']");
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
                        var recordID = $("input[name ='_recordID']");
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
                        var recordID = $("input[name ='_recordID']");
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
    }
});
//产品分类
$("#productCategoryForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
        name: {
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
                        var recordID = $("input[name ='_recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            },
        },
    }
});
//产品属性
$("#productAttributeForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
        name: {
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
                        var recordID = $("input[name ='_recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        }
    }
});
//产品属性值
$("#productAttributeValueForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
        name: {
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
                        var recordID = $("input[name ='_recordID']");
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
        }
    }
});
//计量单位分类
$("#productUomCategForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
        name: {
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
                        var recordID = $("input[name ='_recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        }
    }
});
//计量单位分类
$("#productUomForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
        category: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "计量单位分类不能为空"
                },
            }
        },
        name: {
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
                        var recordID = $("input[name ='_recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
        }
    },
});
//产品款式
$("#productTemplateForm").bootstrapValidator({
    message: '该值无效',
    feedbackIcons: { /*input状态样式图片*/
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
    submitButtons: 'button[type="submit"]',
    trigger: null,
    fields: {
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
                        var recordID = $("input[name ='_recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                },
            },
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
    }
});
// });
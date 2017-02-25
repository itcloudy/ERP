var BootstrapValidator = function(selector, needValidatorFields) {
    //根据数据类型获得正确的数据,默认string
    var getCurrentDataType = function(val, dataType) {
        if (dataType == "" || dataType === undefined || dataType === null) {
            dataType = "string";
        }
        switch (dataType) {
            case "bool":
                val = bool(val);
                break;
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
        }
        return val
    };
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
        // 默认为表单数据创建
        var httpMethod = "POST";
        var $form = $(e.currentTarget);
        var formData = {
            FormAction: "create"
        };
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            xsrf = xsrf[0].value;
            // formData._xsrf = xsrf[0].value;
        }
        //    获得form直接的字段
        var formFields = $form.find(".form-create,.form-edit");
        if ($form.find("input[name='recordID']").length > 0) {
            // 表单数据更新
            formData.FormAction = "update";
            httpMethod = "PUT";
        }
        var actionFields = [];
        for (var item = 0, formFieldsLen = formFields.length; item < formFieldsLen; item++) {
            var self = formFields[item];
            var dataType = $(self).data("type");
            var oldValue = null;
            oldValue = $(self).data("oldvalue");
            var val = $(self).val();
            // 处理radio数据
            if (self.type == "radio") {
                if (dataType == "string") {
                    var nodeName = $("input[name ='" + self.name + "']:checked");
                    if (nodeName != undefined) {
                        actionFields.push(self.name);
                        formData[self.name] = nodeName.val();
                    }
                } else {
                    console.log("data  type is not string");
                }
            } else if (self.type == "checkbox") {
                if (self.checked) {
                    formData[self.name] = true;
                } else {
                    formData[self.name] = false;
                }
                actionFields.push(self.name);
            } else {

                // 判断整形数组值是否改变,只存在增加、删除的情况。oldValue="1,2,3,"
                if (dataType == "array_int") {
                    var addIds = []; //值为记录的id ,int类型
                    var deleteIds = []; //值为记录的id ,int类型
                    console.log(self.name);
                    console.log(oldValue);
                    console.log(val);

                    if (!val) {
                        val = [];
                    }
                    if (!oldValue) {
                        oldValue = "";
                    }

                    var newIdsStr = "," + val.join(",") + ",";
                    var oldValueArrs = oldValue.split(","); //字符分割
                    var oldIdsStr = "," + oldValue;

                    // 如果当前值在旧值中不存在，添加到addIds中
                    for (var nIdex = 0, nLen = val.length; nIdex < nLen; nIdex++) {
                        if (val[nIdex] == "") {
                            continue;
                        }
                        // oldIdsStr =",1,2,3,4," 判断时以","作为分割
                        if (oldIdsStr.indexOf("," + val[nIdex] + ",") == -1) {
                            var newId = parseInt(val[nIdex]);
                            if (newId) {
                                addIds.push(newId);
                            }
                        }
                    }
                    // 如果旧值在当前值中不存在,则认为该记录要被删除
                    for (var oIdex = 0, oLen = oldValueArrs.length; oIdex < oLen; oIdex++) {
                        if (oldValueArrs[oIdex] == "") {
                            continue;
                        }
                        if (newIdsStr.indexOf("," + oldValueArrs[oIdex] + ",") == -1) {
                            var deleteId = parseInt(oldValueArrs[oIdex]);
                            if (deleteId) {
                                deleteIds.push(deleteId);
                            }
                        }
                    }
                    var mapActionRecords = {};
                    if (addIds.length > 0) {
                        mapActionRecords.create = addIds;
                    }
                    if (deleteIds.length > 0) {
                        mapActionRecords.delete = deleteIds;
                    }
                    if (mapActionRecords.delete != undefined || mapActionRecords.create != undefined) {
                        formData[self.name] = mapActionRecords;
                    }
                } else {
                    // 如果值未改变不添加进去
                    if (val == oldValue) {
                        continue;
                    }
                    if (val != "") {
                        // 若为null跳出此次循环
                        if (val === null) {
                            continue;
                        }
                        // 如果input[@name="recordID"]存在
                        if (self.name == 'recordID') {
                            formData["id"] = getCurrentDataType(val, dataType);
                        } else {
                            formData[self.name] = getCurrentDataType(val, dataType);
                            actionFields.push(self.name);
                        }
                    }
                }
            }
        }
        formData["actionFields"] = actionFields;
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
        var requestParams = {
            postData: JSON.stringify(formData),
            _xsrf: xsrf
        };
        var method = $form.find("input[name='_method']");
        if (method.length > 0) {
            requestParams._method = method.val();
        }
        $.ajax({
            type: httpMethod,
            url: $form.action,
            data: requestParams,
            dataType: "json",
            success: function(response) {
                if (response.code == 'failed') {
                    if (formData.FormAction == "update") {
                        toastr.error("修改失败<br>" + response.debug, "错误");
                    } else {
                        toastr.error("创建失败<br>" + response.debug, "错误");
                    }
                    return;
                } else {
                    if (formData.FormAction == "update") {
                        toastr.success("<h3>修改成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
                    } else {
                        toastr.success("<h3>创建成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
                    }
                    console.log(response.location);
                    // setTimeout(function() { window.location = response.location; }, 1000);
                }
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                console.log(XMLHttpRequest.status);
                console.log(XMLHttpRequest.readyState);
                console.log(textStatus);
                toastr.error("请求失败，请刷新页面后再操作", "错误");
            }
        });
        // $.post($form.action, requestParams).success(function(response) {
        //     if (response.code == 'failed') {
        //         if (formData.FormAction == "update") {
        //             toastr.error("修改失败", "错误");
        //         } else {
        //             toastr.error("创建失败", "错误");
        //         }
        //         return;
        //     } else {
        //         if (formData.FormAction == "update") {
        //             toastr.success("<h3>修改成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
        //         } else {
        //             toastr.success("<h3>创建成功</h3><br><a href='" + response.location + "'>1秒后跳转</a>");
        //         }
        //         console.log(response.location);
        //         // setTimeout(function() { window.location = response.location; }, 1000);
        //     }
        // });

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
            }
        },
        NameZh: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "用户名(中文)不能为空"
                }
            }
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
                },
                regexp: {
                    regexp: /^[0-9-]+$/,
                    message: '手机号码只能为数字和 - '
                }
            }
        },
        Tel: {
            message: "该值无效",
            validators: {
                regexp: {
                    regexp: /^[0-9-]+$/,
                    message: '座机只能为数字和 - '
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
                },
                identical: { //相同
                    field: 'ConfirmPassword',
                    message: '两次密码不一致'
                },
                different: { //不能和用户名相同
                    field: 'Name,Email,Mobile',
                    message: '不能和用户名,手机,邮箱,相同'
                },
                regexp: { //匹配规则
                    regexp: /^[a-zA-Z0-9_\.]+$/,
                    message: '密码只能是字母数字下划线'
                }
            }
        },
        ConfirmPassword: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "确认密码不能为空"
                },
                identical: { //相同
                    field: 'Password',
                    message: '两次密码不一致'
                },
                different: { //不能和用户名相同
                    field: 'Name,Email,Mobile',
                    message: '不能和用户名,手机,邮箱,相同'
                },
                different: { //不能和用户名相同
                    field: 'Name,Email,Mobile',
                    message: '不能和用户名,手机,邮箱,相同'
                },
                regexp: { //匹配规则
                    regexp: /^[a-zA-Z0-9_\.]+$/,
                    message: '密码只能是字母数字下划线'
                }
            }
        }
    });
    // 公司
    BootstrapValidator("#companyForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "公司名称不能为空"
                },
                remote: {
                    url: "/company/",
                    message: "公司名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        },
        Code: {
            message: "该值无效",
            validators: {
                remote: {
                    url: "/company/",
                    message: "公司编码已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        },
        Province: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "省份不能为空"
                }
            }
        },
        City: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "城市不能为空"
                }
            }
        },

        Street: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "街道不能为空"
                }
            }
        }
    });
    // 部门
    BootstrapValidator("#departmentForm", {
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "部门名称不能为空"
                },
                remote: {
                    url: "/department/",
                    message: "部门名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {

                        var params = {
                            action: "validator"
                        }

                        var company = $("#company");
                        if (company.length > 0) {
                            company = company.val();
                            if (!company) {
                                toastr.error("请先选择<strong>所属公司</strong>", "错误");
                            }
                            params.company = company;
                        } else {
                            toastr.error("没有公司选项", "错误");
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
            }
        }
    });
    // 团队
    BootstrapValidator("#teamForm", {
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        Leader: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "负责人不能为空"
                },
            }
        },
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "团队名称不能为空"
                },
                remote: {
                    url: "/team/",
                    message: "团队名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        }
    });
    BootstrapValidator("#positionForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "职位名称不能为空"
                },
                remote: {
                    url: "/position/",
                    message: "职位名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        }
    });
    // 菜单
    BootstrapValidator("#menuForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "菜单名称不能为空"
                },
                remote: {
                    url: "/menu/",
                    message: "菜单名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        },
        Identity: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "菜单唯一标识不能为空"
                },
                remote: {
                    url: "/menu/",
                    message: "菜单唯一标识已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        }
    });
    // 角色
    BootstrapValidator("#roleForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "角色名称不能为空"
                },
                remote: {
                    url: "/role/",
                    message: "角色名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        }
    });
    // 资源
    BootstrapValidator("#sourceForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "资源名称不能为空"
                },
                remote: {
                    url: "/source/",
                    message: "资源名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        },
        ModelName: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "Model名称不能为空"
                },
                remote: {
                    url: "/source/",
                    message: "Model名称已经存在",
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
                        var recordID = $("input[name ='recordID']");
                        if (recordID.length > 0) {
                            params.recordID = recordID[0].value;
                        }
                        return params
                    },
                }
            }
        }
    });
    // 合作伙伴
    BootstrapValidator("#partnerForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "名称不能为空"
                },
                remote: {
                    url: "/partner/",
                    message: "名称已经存在",
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
                }
            }
        },
        Mobile: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "手机号码不能为空"
                }
            }
        },
        Province: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "省份不能为空"
                }
            }
        },
        City: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "城市不能为空"
                }
            }
        },

        Street: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "街道不能为空"
                }
            }
        }
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
        }
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
    // 柜台
    BootstrapValidator("#saleCounterForm", {
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "柜台名称不能为空"
                },
                remote: {
                    url: "/sale/counter/",
                    message: "该柜台名称已经存在",
                    dataType: "json",
                    delay: 200,
                    type: "POST",
                    data: function() {
                        var params = {
                            action: "validator",
                        }
                        var company = $("#company");
                        if (company.length > 0) {
                            company = company.val();
                            if (!company) {
                                toastr.error("请先选择<strong>所属公司</strong>", "错误");
                            }
                            params.company = parseInt(company);
                        } else {
                            toastr.error("没有公司选项", "错误");
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
    // 仓库管理
    BootstrapValidator("#stockWarehouseForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "仓库名称不能为空"
                },
                remote: {
                    url: "/stock/warehouse/",
                    message: "仓库名称已经存在",
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
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        Code: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "仓库编码不能为空"
                },
                remote: {
                    url: "/stock/warehouse/",
                    message: "仓库编码已经存在",
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
        Province: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "省份不能为空"
                }
            }
        },
        City: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "城市不能为空"
                }
            }
        },

        Street: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "街道不能为空"
                }
            }
        }
    });
    //库位类型管理
    BootstrapValidator("#stockPickingTypeForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "仓库名称不能为空"
                },
                remote: {
                    url: "/stock/picking/type/",
                    message: "仓库名称已经存在",
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
        WareHouse: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "仓库名称不能为空"
                }
            }
        }
    });
    // 库位
    BootstrapValidator("#stockLocationForm", {
        Name: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "库位名称不能为空"
                },
                remote: {
                    url: "/stock/location/",
                    message: "库位名称已经存在",
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
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        Usage: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "库位类型不能为空"
                },
            }
        },
        Posx: {
            message: "该值无效",
            validators: {
                regexp: {
                    regexp: /[1-9]+$/,
                    message: '通道(X)应为大于0的整数'
                },
            }
        },
        Posy: {
            message: "该值无效",
            validators: {
                regexp: {
                    regexp: /[1-9]+$/,
                    message: '货架(Y)应为大于0的整数'
                },
            }
        },
        Posz: {
            message: "该值无效",
            validators: {
                regexp: {
                    regexp: /[1-9]+$/,
                    message: '层应为大于0的整数'
                },
            }
        }
    });
    // 销售订单
    BootstrapValidator("#saleOrderForm", {
        Partner: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "客户不能为空"
                }
            }
        },
        Company: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "所属公司不能为空"
                },
            }
        },
        StockWarehouse: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "发货仓库不能为空"
                },
            }
        },
        SalesMan: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "业务员不能为空"
                },
            }
        },
        PickingPolicy: {
            message: "该值无效",
            validators: {
                notEmpty: {
                    message: "发货策略不能为空"
                },
            }
        }
    });
});
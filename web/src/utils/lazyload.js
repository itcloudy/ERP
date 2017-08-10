function lazyLoadComponent(path, name) {
    let component;
    try {
        if (path != "") {
            component = resolve => require(['@/page/' + path + '/' + name], resolve);
        } else {
            component = resolve => require(['@/page/' + name], resolve);
        }
    } catch (e) {
        component = resolve => require(['@/page/global/notFound'], resolve);
    }
    return component;
}

function expandItemPath(item) {
    //判断是否存在不同视图
    let viewType = item.ViewType;
    if (viewType.length > 0) {
        let viewTypes = viewType.split(",");
        if (viewTypes.length > 0) {
            item.children = [];
            let routesDict = {};
            for (let i = 0; i < viewTypes.length; i++) {
                let viewType = viewTypes[i].replace(/^\s+|\s+$/g, "");
                if (viewType.length == 0) {
                    continue;
                }
                if ("form" == viewType) {
                    routesDict.form = {};
                    routesDict.form.path = ":id";
                    routesDict.form.component = lazyLoadComponent(item.FloderPath, "Form");

                } else if ('tree' == viewType) {
                    routesDict.tree = {};
                    routesDict.form.path = "/";
                    routesDict.form.component = lazyLoadComponent(item.FloderPath, "Tree");
                }
            }

            for (let key in routesDict) {
                if (!("children" in item)) {
                    item.children = [];
                }
                item.children.push(routesDict[key]);
            }
        }
    }
}
export default function lazyload(menus) {
    console.log(menus);
    //只支持3层菜单，多层请自行采用递归或者嵌套
    menus.map(function(menu) {
        console.log(menu.name);
        console.log(menu);
        if (menu.children != null) {
            menu.expandMenu = true;
            menu.component = lazyLoadComponent("global", "Home");
            menu.children.map(function(item) {
                if (item.children != null) {
                    item.expandMenu = true;
                    item.component = lazyLoadComponent(item.FloderPath, item.Component.replace(/^\s+|\s+$/g, ""));
                    item.children.map(function(su) {
                        // 不再支持更深菜单
                        su.expandMenu = false;
                        su.component = lazyLoadComponent(su.FloderPath, su.Component.replace(/^\s+|\s+$/g, ""));
                        if (su.viewTypePaths) {
                            let paths = su.viewTypePaths;
                            su.children = [];
                            for (let i = 0; i < paths.length; i++) {
                                let pathItem = paths[i];
                                let itemCom = {};
                                itemCom.path = pathItem.path;
                                itemCom.component = lazyLoadComponent(su.FloderPath, pathItem.Component);
                                item.children.push(itemCom);
                            }
                        }
                    });
                } else {
                    item.expandMenu = false;
                    item.component = lazyLoadComponent(item.FloderPath, item.Component.replace(/^\s+|\s+$/g, ""));
                    if (item.viewTypePaths) {
                        let paths = item.viewTypePaths;
                        item.children = [];
                        for (let i = 0; i < paths.length; i++) {
                            let pathItem = paths[i];
                            let itemCom = {};
                            itemCom.path = pathItem.path;
                            itemCom.component = lazyLoadComponent(item.FloderPath, pathItem.Component);
                            item.children.push(itemCom);
                        }
                    }
                }
            });
        } else {
            menu.expandMenu = false;
            menu.component = lazyLoadComponent(menu.FloderPath, menu.Component.replace(/^\s+|\s+$/g, ""));
            if (menu.viewTypePaths) {
                let paths = menu.viewTypePaths;
                menu.children = [];
                for (let i = 0; i < paths.length; i++) {
                    let pathItem = paths[i];
                    let itemCom = {};
                    itemCom.path = pathItem.path;
                    itemCom.component = lazyLoadComponent(menu.FloderPath, pathItem.Component);
                    menu.children.push(itemCom);
                }
            }
        }
    });
    return menus;
}
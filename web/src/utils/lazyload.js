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

export default function lazyload(menus) {
    //只支持3层菜单，多层请自行采用递归或者嵌套
    menus.map(function(menu) {
        if (menu.children != null) {
            menu.component = lazyLoadComponent("global", "Home");
            menu.children.map(function(item) {
                if (item.children != null) {
                    item.component = lazyLoadComponent(item.FloderPath, item.Component.replace(/^\s+|\s+$/g, ""));
                    item.children.map(function(su) {
                        不再支持更深菜单
                        if (su.children != null) {
                            su.component = lazyLoadComponent(su.FloderPath, su.Component.replace(/^\s+|\s+$/g, ""));
                        } else {
                            su.component = lazyLoadComponent(su.FloderPath, su.Component.replace(/^\s+|\s+$/g, ""));
                        }

                    });
                } else {
                    item.component = lazyLoadComponent(item.FloderPath, item.Component.replace(/^\s+|\s+$/g, ""));
                    //判断是否存在不同视图
                    // let components = item.Component;
                    // components = components.split(",");
                    // let len = components.length;
                    // if (len == 1) {
                    //     item.component = lazyLoadComponent(item.FloderPath, item.Component.replace(/^\s+|\s+$/g, ""));
                    // } else {
                    //     item.children = [];
                    //     for (let i = 0; i < len; i++) {
                    //         let type = components[i];
                    //         type = type.replace(/^\s+|\s+$/g, "");
                    //         if ("form" == type) {
                    //             item.children.push({
                    //                 path: "/:id",
                    //                 component: lazyLoadComponent(item.FloderPath, "Form"),
                    //             });
                    //         } else {
                    //             item.children.push({
                    //                 path: "/",
                    //                 component: lazyLoadComponent(item.FloderPath, "Tree"),
                    //             });
                    //         }
                    //     }

                    // }
                }
            });
        } else {
            menu.component = lazyLoadComponent(menu.FloderPath, menu.Component.replace(/^\s+|\s+$/g, ""));
        }
    });
    return menus;
}
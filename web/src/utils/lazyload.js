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
                    item.component = lazyLoadComponent("global", "Home");
                    item.children.map(function(su) {
                        if (su.children != null) {
                            // su.component = lazyload("global", "Home");
                        } else {
                            su.component = lazyLoadComponent(su.Category, su.Component);
                        }
                    });
                } else {
                    item.component = lazyLoadComponent(item.Category, item.Component);
                }
            });
        } else {
            menu.component = lazyLoadComponent(menu.Category, menu.Component);
        }
    });
    return menus;
}
function lazyLoadComponent(path) {
    return resolve => require(['@/views/' + path], resolve);
}


export default function lazyLoadMenusRoutes(menus) {

    menus.map(function(menu) {
        if (menu.children != null) {
            menu.expand = true;
            menu.component = lazyLoadComponent(menu.ComponentPath);
            menu.children.map(function(item) {
                if (item.children != null) {
                    item.expand = true;
                    item.component = lazyLoadComponent(item.ComponentPath.replace(/^\s+|\s+$/g, ""));
                    item.children.map(function(su) {
                        // 不再进一步扩展
                        su.component = lazyLoadComponent(su.ComponentPath.replace(/^\s+|\s+$/g, ""));
                        su.expand = false;
                        let viewType = su.ViewType;
                        if (viewType) {
                            let viewTypeJsons = JSON.parse(viewType);
                            for (let i = 0; i < viewTypeJsons.length; i++) {
                                let n = viewTypeJsons[i];
                                if (i == 0) {
                                    su.children = [];
                                }
                                su.children.push({
                                    path: "/" + n.path,
                                    component: lazyLoadComponent(n.componentpath),
                                })
                            }
                        }
                    });
                } else {
                    item.component = lazyLoadComponent(item.ComponentPath.replace(/^\s+|\s+$/g, ""));
                    item.expand = false;
                    let viewType = item.ViewType;
                    if (viewType) {
                        let viewTypeJsons = JSON.parse(viewType);
                        for (let i = 0; i < viewTypeJsons.length; i++) {
                            let n = viewTypeJsons[i];
                            if (i == 0) {
                                item.children = [];
                            }
                            item.children.push({
                                path: "/" + n.path,
                                component: lazyLoadComponent(n.componentpath),
                            })
                        }
                    }
                }
            });
        } else {
            menu.component = lazyLoadComponent(menu.ComponentPath.replace(/^\s+|\s+$/g, ""));
            menu.expand = false;
            let viewType = menusu.ViewType;
            if (viewType) {
                let viewTypeJsons = JSON.parse(viewType);
                for (let i = 0; i < viewTypeJsons.length; i++) {
                    let n = viewTypeJsons[i];
                    if (i == 0) {
                        menu.children = [];
                    }
                    menu.children.push({
                        path: "/" + n.path,
                        component: lazyLoadComponent(n.componentpath),
                    })
                }
            }
        }
    });
    //外层增加Home
    let bootMenu = {
        path: "",
        component: lazyLoadComponent("global/Home"),
        children: menus
    }
    return bootMenu;
}
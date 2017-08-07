import * as types from './mutations-types'
import { lazyload } from '../utils/lazyload'

export default {
    [types.GLOBAL_SET_WINDOW_HRIGHT](state, height) {
        state.windowStyles.leftSidebarStyles.height = height + 'px';
        let mainHeight = height - 50;
        state.windowStyles.rightMainStyles.height = mainHeight + 'px';
    },
    [types.GLOBAL_TOGGLE_LEFT_SIDEBAR](state) {
        if (state.showLeftSidebar)
            state.showLeftSidebar = false;
        else
            state.showLeftSidebar = true;
    },
    [types.GLOBAL_SET_USERINFO](state, userinfo) {
        state.userinfo = userinfo;
    },
    [types.GLOBAL_SET_UER_MENUS](state, menus) {
        console.log(menus);
        //只支持3层菜单，多层请自行采用递归或者嵌套
        menus.map(function(menu) {
            if (menu.children != null) {
                menu.component = lazyload("global", "Home");
                menu.children.map(function(item) {
                    if (item.children != null) {
                        item.component = lazyload("global", "Home");
                        item.children.map(function(su) {
                            if (su.children != null) {
                                // su.component = lazyload("global", "Home");
                            } else {
                                su.component = lazyload(su.Category, su.Component);
                            }
                        });
                    } else {
                        item.component = lazyload(item.Category, item.Component);
                    }
                });
            } else {
                menu.component = lazyload(menu.Category, menu.Component);
            }
        });
        console.log(menus);
        state.menus = menus;
    },
    [types.GLOBAL_LOAD_ROUTES_DONE](state, done) {
        state.loadRoutersDone = true;
    }

}
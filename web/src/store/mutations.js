import * as types from './mutations-types'

export default {
    [types.GLOBAL_SET_WINDOW_HRIGHT](state, height) {
        state.windowStyles.leftSidebarStyles.height = height + 'px';
        let mainHeight = height - 50;
        state.windowHeight = mainHeight;
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
        state.isBackgroundUser = userinfo.isBackground;
    },
    [types.GLOBAL_SET_UER_MENUS](state, menus) {
        state.backgroundMenus = menus;
    },
    [types.GLOBAL_LOAD_ROUTES_DONE](state, done) {
        state.loadRoutersDone = true;
    }

}
import * as types from './mutations-types'

export default {
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
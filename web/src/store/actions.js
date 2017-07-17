import * as types from './mutations-types'

export default {
    GLOBAL_TOGGLE_LEFT_SIDEBAR: ({ commit }) => {
        commit(types.GLOBAL_TOGGLE_LEFT_SIDEBAR)
    },
    GLOBAL_SET_WINDOW_HRIGHT: ({ commit }, { height }) => {
        commit(types.GLOBAL_SET_WINDOW_HRIGHT, height)
    }


}
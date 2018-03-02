import Vue from 'vue'
import Vuex from 'vuex'

// 根级state、actions、getters、mutations
import state from './state';
import actions from './actions';
import getters from './getters';
import mutations from './mutations';
// 合作伙伴信息
import adminPartner from './partner';
// 产品管理
import adminProduct from './product';
// 采购管理
import adminPurchase from './purchase';
// 销售管理
import adminSale from './sale';
// 库存管理
import adminStock from './stock'
// 用户管理
import adminUser from './user';



Vue.use(Vuex)
const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    state,
    actions,
    getters,
    mutations,
    modules: {
        adminPartner, //合作伙伴
        adminProduct, //产品
        adminPurchase, //采购
        adminSale, //销售
        adminStock, //库存
        adminUser, //用户
    },
    strict: debug,
})
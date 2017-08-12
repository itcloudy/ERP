import {
    db_prefix
} from '../config/settings';

class LocalStore {
    constructor() {
        this.localStore = window.localStorage;
        this.prefix = db_prefix;
    }
    set(key, value, fn) {
        try {
            value = JSON.stringify(value);
        } catch (e) {
            value = value;
        }

        this.localStore.setItem(this.prefix + key, value);

        fn && fn();
    }
    get(key, fn) {
        if (!key) {
            throw new Error('没有找到key。');
            return;
        }
        if (typeof key === 'object') {
            throw new Error('key不能是一个对象。');
            return;
        }
        var value = this.localStore.getItem(this.prefix + key);
        if (value !== null) {
            try {
                value = JSON.parse(value);
            } catch (e) {
                value = value;
            }
        }

        return value;
    }
    remove(key) {
        this.localStore.removeItem(this.prefix + key);
    }
}
const localStore = new LocalStore();
export default localStore
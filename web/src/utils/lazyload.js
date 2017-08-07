export function lazyload(path, name) {
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
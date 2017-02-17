$(function() {
    // 'table-product-product'
    //对产品规格数据进行处理
    $.contextMenu({
        selector: '#table-product-product ',
        build: function($trigger, e) {
            return {
                callback: function(key, options) {
                    var m = "clicked: " + key;
                    console.log(m);
                },
                items: {
                    "edit": { name: "Edit", icon: "edit" },
                    "cut": { name: "Cut", icon: "cut" },
                    "status": {
                        name: "Status",
                        icon: "delete",
                        items: loadItems(),
                    },
                    "normalSub": {
                        name: "Normal Sub",
                        items: {
                            "normalsub1": { name: "normal Sub 1" },
                            "normalsub2": { name: "normal Sub 2" },
                            "normalsub3": { name: "normal Sub 3" },
                        }
                    }
                }
            };
        }
    });
});
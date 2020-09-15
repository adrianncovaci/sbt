import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import PolicyExport from './components/PolicyExport.vue';
import FileParser from './components/FileParser.vue';

export default new Router({
    routes: [{
        path: "/export",
        name: "policy-export",
        component: PolicyExport
    },
    {
        path: "",
        name: "file-parser",
        component: FileParser
    }
    ]
});
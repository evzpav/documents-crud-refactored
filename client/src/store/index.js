import Vue from 'vue';
import Vuex from 'vuex';
import axios from "axios"
import {endpoints} from "../configs/global";

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        isDataReady: false
    },
    getters: {
        isDataReady: (state) => state.isDataReady
    },
    mutations: {
        setDataReady: (state, payload) => {
            state.isDataReady = payload
        },
    },
    actions: {
        listDocuments(context) {
            return baseRequest(context, "get", "/documents")
        },
        deleteDocument(context, id) {
            return baseRequest(context, "delete", `/document/${id}`)
        },
        createDocument(context, form) {
            return baseRequest(context, "post", "/document", form)
        },
        updateDocument(context, form) {
            return baseRequest(context, "put", `/document/${form.id}`, form)
        },
        serverStatus(context) {
            return baseRequest(context, "get", `/status`)
        },



    }
});

async function baseRequest(context, method, uri, data) {
    context.commit("setDataReady", false);
    try {
        const url = endpoints.server + uri;
        let headers = {};
        const resp = await axios({
            method: method,
            url: url,
            data: data,
            headers: headers,
        });
        return resp
    } catch (error) {
        if (error.response) {
            if (error.response.data) {
                throw new Error(error.response.data.message)
            }
        } else if (error.request) {
            if (!error.status) {
                console.log(error.message)
            }
        } else {
            throw new Error(error.message);
        }
        throw error;
    }

}

export default store;

import Documents from '../components/Documents.vue'
import Status from '../components/Status.vue'

export default [
    {
        path: '/',
        component: Documents
    },
    {path: '/status', name: 'status', component: Status},
    // {
    //     path: '*', component: Page404,
    // },

]

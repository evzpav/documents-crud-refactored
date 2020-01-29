<template>
    <div class="column container_sm">
        <div class="body_height">
            <div class="column">
                <h1 class="title">Documents CRUD</h1>
            </div>

            <div class="columns column" id="new_document_btn_div">
                <div class="column">
                    <div class="control search_input">
                        <the-mask type="text" class="input" placeholder="Search CPF/CNPJ"
                                  v-model="searchInput" :mask="['###.###.###-##', '##.###.###/####-##']">
                        </the-mask>
                    </div>
                </div>
                <div class="field column pb0">
                    <div class="control filter_doc_type">
                        <b-select placeholder="Type" v-model="filterDocType">
                            <option value="">All Types</option>
                            <option value="CPF">CPF</option>
                            <option value="CNPJ">CNPJ</option>
                        </b-select>
                    </div>
                </div>
                <div class="field column pb0 filter_is_blacklisted">
                    <div class="control">
                        <b-select placeholder="Type" v-model="filterIsBlacklisted">
                            <option value="">Show all</option>
                            <option :value="true">Blacklisted</option>
                            <option :value="false">Not blacklisted</option>
                        </b-select>
                    </div>
                </div>
                <div class="field column new_doc_btn">
                    <button @click="cleanFilters()" class="button">Clean Filters</button>
                </div>
                <div class="field column new_doc_btn">
                    <button @click="openDocumentModal()" class="button success">New Document</button>
                </div>

            </div>
            <doc-modal></doc-modal>
            <div class="column" v-if="!isLoading">

                <vue-good-table
                        :columns="columns"
                        :rows="filteredDocuments"
                        :sort-options="{
                    enabled: true,
                  }"
                        :pagination-options="{
                    enabled: true,
                    mode: 'records',
                    perPage: 10,
                    position: 'bottom',
                    dropdownAllowAll: false,
                    setCurrentPage: 1,
                    nextLabel: 'next',
                    prevLabel: 'prev',
                    rowsPerPageLabel: 'Rows per page',
                    ofLabel: 'of',
                    pageLabel: 'page', // for 'pages' mode
                    allLabel: 'All',
                  }"
                >
                    <template slot="table-row" slot-scope="props">
            <span v-if="props.column.field === 'value'">
              <the-mask
                      :mask="['###.###.###-##', '##.###.###/####-##']"
                      :masked="true"
                      :value="props.row.value"
                      class="control table_mask"
                      disabled
              ></the-mask>
            </span>
                        <span v-if="props.column.field === 'isBlacklisted'">
              <span v-if="props.row.isBlacklisted === true">Yes</span>
              <span v-else>No</span>
            </span>
                        <span v-if="props.column.field === 'actions'">
              <button @click="openDocumentModal(props.row)" class="action_button">
                <span v-html="icons.edit" class="normal_icon"></span>
              </button>
              <button @click="confirmCustomDelete(props.row.id)" class="action_button">
                <span v-html="icons.delete" class="normal_icon"></span>
              </button>
            </span>
                        <span
                                v-if="props.column.field !== 'value' && props.column.field !== 'actions' && props.column.field !== 'isBlacklisted'"
                        >{{props.formattedRow[props.column.field]}}</span>
                    </template>
                </vue-good-table>
            </div>
            <div v-else>Loading documents...</div>
        </div>

        <div>
            <the-footer></the-footer>
        </div>
    </div>
</template>

<script>
    import TheFooter from "./TheFooter.vue";
    import DocModal from "./DocModal.vue";
    import {icons} from "../icons.js";
    import {mapGetters} from "vuex";

    function Form() {
        this.isBlacklisted = false;
        this.docType = "CPF";
        this.value = "";
        this.id = "";
    }

    export default {
        components: {
            TheFooter,
            DocModal
        },
        mounted() {
            this.listDocuments();

        },
        data: () => ({
            filteredDocuments: [],
            filterIsBlacklisted: "",
            icons,
            filterDocType: "",
            searchInput: "",
            form: new Form(),
            isLoading: false,
            documents: [],
            columns: [
                {
                    label: "Document",
                    field: "value",
                    type: "string"
                },
                {
                    label: "Type",
                    field: "docType",
                    type: "string"
                },
                {
                    label: "Blacklisted",
                    field: "isBlacklisted",
                    type: "boolean"
                },
                {
                    label: "Actions",
                    field: "actions"
                }
            ]
        }),

        methods: {
            async listDocuments() {
                this.isLoading = true;
                try {
                    const documents = await this.$store.dispatch("listDocuments");
                    this.documents = documents.data;
                    this.filteredDocuments = this.documentsFilterSearch()
                } catch (e) {
                    console.log(e);
                } finally {
                    this.isLoading = false;
                }
            },
            async deleteDocument(id) {
                this.isLoading = true;
                try {
                    await this.$store.dispatch("deleteDocument", id);
                } catch (e) {
                    console.log(e);
                } finally {
                    this.isLoading = false;
                    this.$store.commit("setDataReady", true)
                }
            },
            openDocumentModal(docToEdit) {
                if (docToEdit) {
                    docToEdit["edit"] = true;
                }
                this.$modal.show("new-document", docToEdit);
            },
            confirmCustomDelete(id) {
                this.$dialog.confirm({
                    title: "Deleting document",
                    message: "Are you sure you want to <b>delete</b> this document?",
                    confirmText: "Delete document",
                    type: "is-danger",
                    hasIcon: true,
                    onConfirm: () => {
                        this.showToast("Document deleted", false);

                        this.deleteDocument(id);
                    }
                });
            },

            showToast(msg, isSuccess) {
                this.$toast.open({
                    message: msg,
                    type: isSuccess ? "is-success" : "is-danger",
                    position: "is-bottom"
                });
            },
            searchString(item, input) {
                if (input) {
                    item = item + "";
                    return item.toLowerCase().includes(input.toLowerCase());
                } else {
                    return false;
                }
            },
            documentsFilterSearch() {
                const vm = this;
                if(this.documents) {
                    return this.documents.filter(doc => {
                        let docSearchFiltered = false;
                        let docTypeFiltered = false;
                        let isBlacklistedFiltered = false;

                        if (vm.searchInput) {
                            docSearchFiltered = vm.searchString(doc.value, vm.searchInput)
                        } else {
                            docSearchFiltered = true
                        }

                        if (vm.filterDocType) {
                            docTypeFiltered = doc.docType === vm.filterDocType;
                        } else {
                            docTypeFiltered = true
                        }

                        if (vm.filterIsBlacklisted !== "" && vm.filterIsBlacklisted !== undefined) {
                            isBlacklistedFiltered = doc.isBlacklisted === vm.filterIsBlacklisted;
                        } else {
                            isBlacklistedFiltered = true;
                        }

                        return docSearchFiltered && docTypeFiltered && isBlacklistedFiltered;

                    });
                } else{
                    return [];
                }
            },
            cleanFilters() {
                this.searchInput = "";
                this.filterIsBlacklisted = "";
                this.filterDocType = "";
            }
        },
        computed: {
            ...mapGetters(["isDataReady"])
        },
        watch: {
            isDataReady() {
                if (this.isDataReady) {
                    this.listDocuments();
                    this.filteredDocuments = this.documentsFilterSearch();
                }
            },
            searchInput: function (newValue, old) {

                this.filteredDocuments = this.documentsFilterSearch();
            },
            filterDocType: function (newValue, old) {
                this.filteredDocuments = this.documentsFilterSearch();
            },
            filterIsBlacklisted: function (newValue, old) {
                this.filteredDocuments = this.documentsFilterSearch();
            },


        }
    };
</script>



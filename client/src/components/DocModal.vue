<template>
    <modal name="new-document" class="custom_modal" @before-open="beforeOpen">
        <div class="column">
            <div class="column">
                <div class="column" v-if="!modalForm.edit">
                    <h1 class="title custom_title">New document</h1>
                </div>
                <div class="column" v-else>
                    <h1 class="title custom_title">Edit document</h1>
                </div>

                <div class="field column" label="Document">
                    <div class="control">
                        <the-mask
                                class="input"
                                :mask="['###.###.###-##', '##.###.###/####-##']"
                                v-model="modalForm.value"
                                placeholder="868.335.241-22"
                        />
                    </div>
                </div>
                <div class="columns column is-mobile pb0 custom_select">
                    <div class="field column pb0">
                        <div class="control">
                            <b-select placeholder="Type" v-model="modalForm.docType">
                                <option value="CPF">CPF</option>
                                <option value="CNPJ">CNPJ</option>
                            </b-select>
                        </div>
                    </div>
                    <div class="field column pb0 custom_checkbox">
                        <div class="control">
                            <b-checkbox v-model="modalForm.isBlacklisted">Blacklisted</b-checkbox>
                        </div>
                    </div>
                </div>
                <div class="column pt0">
                    <button
                            class="button success submit_button"
                            @click="submitDocumentForm(modalForm.edit)"
                    >
                        <span v-if="!isLoading">Submit</span>
                        <span v-else>Please wait...</span>
                    </button>
                </div>
            </div>
        </div>
    </modal>

</template>
<script>
    function Form() {
        this.isBlacklisted = false;
        this.docType = "CPF";
        this.value = "";
        this.id = "";
    }

    export default {
        props: ["form"],
        data: () => ({
            isLoading: false,
            modalForm: new Form()
        }),
        methods: {
            beforeOpen(event) {
                this.modalForm = event.params ? event.params : new Form();
            },
            submitDocumentForm(edit) {
                if (!this.modalForm.value || this.modalForm.value.length < 11) {
                    this.showToast("Invalid document. Minimum 11 characters.");
                    return;
                } else if (!this.modalForm.docType) {
                    this.showToast("Invalid document type");
                    return;
                }

                if (!edit) {
                    this.createDocument(this.modalForm);
                } else {
                    this.updateDocument(this.modalForm);
                }

            },
            async createDocument(payload) {
                this.isLoading = true;
                try {
                    await this.$store.dispatch("createDocument", payload);
                } catch (e) {
                    console.log(e);
                } finally {
                    this.hideModal();
                    this.isLoading = false;
                    this.showToast("Document created", true);
                    this.$store.commit("setDataReady", true)
                }
            },
            async updateDocument(payload) {
                this.isLoading = true;
                try {
                    await this.$store.dispatch("updateDocument", payload);
                } catch (e) {
                    console.log(e);
                } finally {
                    this.hideModal();
                    this.isLoading = false;

                    this.showToast("Document updated successfully", true);
                    this.$store.commit("setDataReady", true)
                }
            },
            hideModal() {
                this.$modal.hide("new-document");
            },
            showToast(msg, isSuccess) {
                this.$toast.open({
                    message: msg,
                    type: isSuccess ? 'is-success' : "is-danger",
                    position: "is-bottom"
                });
            }
        },
        watch: {
            "modalForm.value": function (newValue, old) {
                if (newValue.length > 11) {
                    this.modalForm.docType = "CNPJ"
                } else {
                    this.modalForm.docType = "CPF"
                }

            }
        }
    }
</script>

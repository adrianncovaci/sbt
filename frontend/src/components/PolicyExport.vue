<template>
  <div class="container">
    <b-card style="border:none" bg-variant="white" text-variant="dark">
      <b-button to="/" router-tag="li">Import Policies</b-button>
      <hr class="my-4" />
      <b-card-text>Choose Policies:</b-card-text>
      <b-form-checkbox v-model="selectAllPolicies">Select All</b-form-checkbox>
      <hr class="my-4" />
      <b-form-input v-model="selectPolicy" list="policies-title"></b-form-input>
      <datalist id="policies-title">
        <option v-for="title in policiesTitle" :key="title.id">{{title.description}} || {{title.id}}</option>
      </datalist>
      <hr class="my-4" />
      <div id="checkboxes">
        <b-form-group v-model="selectPolicy" label="Policies:">
          <v-layout row wrap>
            <v-flex v-for="policy in selectedPolicies" :key="policy.id" xs6>
              <b-form-checkbox :value="policy" :checked="true">{{policy.description}}</b-form-checkbox>
            </v-flex>
          </v-layout>
        </b-form-group>
      </div>
      <hr class="my-4" />
      <b-form-input v-model="fileName" placeholder="Enter File Name"></b-form-input>
      <hr class="my-4" />
      <b-button @click="exportSelectedPolicies">Export</b-button>
    </b-card>
  </div>
</template>

<script>
export default {
  async created() {
    await this.getAllPolicies();
    this.policies.forEach(element => {
      this.policiesTitle.push(Object.assign({}, element));
    });
    this.policiesTitle = this.policiesTitle.sort((a, b) =>
      a.description > b.description ? 1 : -1
    );
  },
  data() {
    return {
      policies: [],
      policiesTitle: [],
      selectedValue: "",
      selectedPolicies: [],
      fileName: ""
    };
  },
  computed: {
    selectAllPolicies: {
      get: function() {
        return this.policies
          ? this.selectedPolicies.length == this.policies.length
          : false;
      },
      set: function(check) {
        var selected = [];
        if (check) {
          this.policies.forEach(element => {
            selected.push(element);
          });
        }
        this.selectedPolicies = selected;
      }
    },
    selectPolicy: {
      get: function() {
        return this.selectedValue;
      },
      set: function(val) {
        this.selectedValue = val;
        var policyId = this.selectedValue.split("||");
        if (policyId[1]) {
          policyId = policyId[1].trim();
          const policy = this.policies.find(x => x.id == policyId);
          if (this.selectedPolicies.includes(policy)) {
            return;
          } else {
            this.selectedPolicies.push(policy);
          }
        }
      }
    }
  },
  methods: {
    async getAllPolicies() {
      await window.backend.getAllPolicies().then(res => {
        this.policies = res.Value;
      });
    },
    async exportSelectedPolicies() {
      const policiesString = JSON.stringify(this.selectedPolicies);
      const policiesResult = policiesString.replace(/},/g, "},\n");
      await window.backend
        .exportPolicies(policiesResult, this.fileName)
        .then(console.log(policiesResult));
    }
  }
};
</script>

<style scoped></style>
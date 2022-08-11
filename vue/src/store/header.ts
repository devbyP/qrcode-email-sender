import { defineStore } from "pinia";

export const useHeaderStore = defineStore("headerState", {
  state: () => ({
    hasBack: false,
    toScan: false,
  }),
  actions: {
    grantBack() {
      this.hasBack = true;
    },
    revokeBack() {
      this.hasBack = false;
    },
    grantScan() {
      this.toScan = true;
    },
    revokeScan() {
      this.toScan = false;
    },
  },
});

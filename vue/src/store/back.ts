import { defineStore } from "pinia";

export const useStore = defineStore("back", {
  state: () => ({
    hasBack: false,
  }),
  actions: {
    grantBack() {
      this.hasBack = true;
    },
    revokeBack() {
      this.hasBack = false;
    },
  },
});

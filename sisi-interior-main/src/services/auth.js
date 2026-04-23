import API from "./api";

export const login = async (data) => {
  return API.post("/login", data);
};
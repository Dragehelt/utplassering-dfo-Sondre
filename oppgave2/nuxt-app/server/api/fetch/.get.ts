export default defineEventHandler(async (event) => {
  interface DuckResponse {
    url: string;
  }

  const data = await $fetch<DuckResponse>("https://random-d.uk/api/random");
  return data.url;
});

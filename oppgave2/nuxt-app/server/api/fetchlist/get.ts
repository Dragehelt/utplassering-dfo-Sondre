export default defineEventHandler(async (event) => {
  interface DuckResponse {
    images: [
        name: string
    ];
    image_count: number;
  }

  const data = await $fetch<DuckResponse>("https://random-d.uk/api/list");
  //   console.log(data);
  return data.images[0];
  //
});

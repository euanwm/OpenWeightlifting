import fetchEventResult from './fetchEventResult';

describe('fetchEventResult', () => {
  it('should return a valid event result', async () => {
    const eventMetaData = {
      "name": "Leangen AK Seriestevne",
      "federation": "NVF",
      "date": "2023-01-07",
      "id": "58.csv"
    }
    const result = await fetchEventResult(eventMetaData);
    expect(result).toEqual(
      {"size":8,"data":[{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Women's Senior 64Kg","lifter_name":"Celine Dorothea Opdal","bodyweight":64,"snatch_1":60,"snatch_2":-62,"snatch_3":0,"cj_1":78,"cj_2":81,"cj_3":0,"best_snatch":60,"best_cj":81,"total":141,"sinclair":183.05206,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Women's Senior 76Kg","lifter_name":"Live Wahl Gellein","bodyweight":76,"snatch_1":-67,"snatch_2":-67,"snatch_3":-67,"cj_1":85,"cj_2":-86,"cj_3":-86,"best_snatch":0,"best_cj":85,"total":0,"sinclair":0,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Women's Senior 71Kg","lifter_name":"Signe Høstmark","bodyweight":71,"snatch_1":63,"snatch_2":65,"snatch_3":-66,"cj_1":85,"cj_2":84,"cj_3":-88,"best_snatch":65,"best_cj":84,"total":149,"sinclair":182.50119,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Women's Senior 76Kg","lifter_name":"Nadine Ohla","bodyweight":76,"snatch_1":65,"snatch_2":69,"snatch_3":-71,"cj_1":85,"cj_2":87,"cj_3":-90,"best_snatch":69,"best_cj":87,"total":156,"sinclair":184.65465,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Men's Senior 81Kg","lifter_name":"Vegard Vikane","bodyweight":81,"snatch_1":-97,"snatch_2":-97,"snatch_3":-97,"cj_1":115,"cj_2":-118,"cj_3":-118,"best_snatch":0,"best_cj":115,"total":0,"sinclair":0,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Men's Senior 96Kg","lifter_name":"Peder Lindsetmo","bodyweight":96,"snatch_1":-105,"snatch_2":-105,"snatch_3":-105,"cj_1":-120,"cj_2":-128,"cj_3":-131,"best_snatch":0,"best_cj":0,"total":0,"sinclair":0,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Men's Senior 96Kg","lifter_name":"Thomas Malmo","bodyweight":96,"snatch_1":100,"snatch_2":105,"snatch_3":110,"cj_1":130,"cj_2":0,"cj_3":0,"best_snatch":110,"best_cj":130,"total":240,"sinclair":270.29492,"country":"NVF","instagram":""},{"event":"Leangen AK Seriestevne","date":"2023-01-07","gender":"Women's Senior 71Kg","lifter_name":"Cecilie Tomassen","bodyweight":71,"snatch_1":64,"snatch_2":66,"snatch_3":-68,"cj_1":81,"cj_2":83,"cj_3":85,"best_snatch":66,"best_cj":85,"total":151,"sinclair":184.95087,"country":"NVF","instagram":""}]}
    )});
});

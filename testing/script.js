import {check} from "k6";
import http from "k6/http";

export const options = {
    scenarios: {
        my_scenario1: {
            executor: 'constant-arrival-rate',
            duration: '30s', // total duration
            preAllocatedVUs: 1000, // to allocate runtime resources

            rate: 1000, // number of constant iterations given `timeUnit`
            timeUnit: '1s',
        },
    },
};

export default function () {
    const url = "http://localhost:4001/api/v1/links";

    // Create Link
    const payload = JSON.stringify({url: "https://github.com/aryahmph"});
    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };
    const resPost = http.post(url, payload, params);
    check(resPost, {
        "Create link status must 201": (r) => r.status === 201,
    });
    const id = JSON.parse(resPost.body).data.id;

    // Get Link 10 times
    for (let i = 0; i < 10; i++) {
        const resGet = http.get(url + `/${id}`);
        check(resGet, {
            "Get link status must 200": (r) => r.status === 200,
            "Get link response id must same with id param": (r) =>
                r.body.includes(id),
        });
    }
}

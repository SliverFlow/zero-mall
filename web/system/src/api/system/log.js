import service from "../../utils/request.js";

export const computerLogApi = () => {
    return service({
        url: "/system/v1/log/computer",
        method: 'post',
    })
}
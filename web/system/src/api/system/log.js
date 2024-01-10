import service from "../../utils/request.js";

export const computerLogApi = () => {
    return service({
        url: "/system/v1/log/computer",
        method: 'post',
    })
}

export const pageListLogApi = (data) => {
    return service({
        url: "/system/v1/log/list",
        method: 'post',
        data: data
    })
}
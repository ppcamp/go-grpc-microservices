
function with_deadline(seconds: number, data?: {}): {} {
    if (!data) data = {};

    var deadline = new Date();
    deadline.setSeconds(deadline.getSeconds() + seconds);
    data["deadline"] = deadline.getTime();

    return data;
}


export default {
    with_deadline
}
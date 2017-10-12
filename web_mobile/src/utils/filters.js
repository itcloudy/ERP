export default function stringTimeFormat(timeString) {
    let date = new Date(timeString);
    let time = date.toLocaleString();
    time = time.replace(/\//g, "-");
    time = time.replace("上午", "");
    time = time.replace("下午", "");
    return time;
}
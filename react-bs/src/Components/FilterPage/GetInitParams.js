const seperate = (str) => {
    const obj = {}
    if(str === ""){
        return obj
    }
    const strArr = str.slice(1).split("&")
    strArr.forEach( (param) => {
        obj[param.split("=")[0]] = param.split("=")[1] 
    })
    return obj
}


const GetInitParams = (query, queryParams) => {
    if (queryParams !== "") {
        const queriesObj = seperate(queryParams)
        for(let key in queriesObj){
            if(key === query){
                return queriesObj[query]
            }
        }
    }
    if (query === "genre") {
        return "All"
    }
    return ""
}

export default GetInitParams;

import DNE from '../../images/Poster_not_available.jpg'


const ProcessImage = (imglink, imageSize) => {
    const arr = imglink.split("/")
    if(arr.length !== 7){
        return DNE
    }
    arr[5] = imageSize
    imglink =  arr.join('/')
    return imglink
}
export default ProcessImage
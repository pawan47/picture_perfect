import React from 'react'
import Pagination from 'react-bootstrap/Pagination'



const Paginate = ({totalMovies, currentPage, moviePerPage, changePage}) => {
    const items = []
    if (totalMovies === 0) {
        return null;
    }
    let totalPage = Math.ceil(totalMovies / moviePerPage)
    const pagelist = []
    for (let i = Math.max(1, currentPage - 2); i <= Math.min(totalPage, currentPage + 2); i++) {
        pagelist.push(i)
    }
    items.push(<Pagination.First key={0} onClick={() => changePage(1)} />)
    items.push(<Pagination.Prev key={-1} disabled={currentPage === 1} onClick={() => {
        if (currentPage === 1) {
            return false;
        }
        changePage(currentPage - 1)
    }} />)
    pagelist.forEach(element => {
        items.push(<Pagination.Item key={element} active={currentPage === element} onClick={() => changePage(element)}>
            {element}
        </Pagination.Item>)
    });
    items.push(<Pagination.Next key = {totalPage + 1} disabled={currentPage === totalPage} onClick={() => {
        if (totalPage === currentPage) {
            return false;
        }
        changePage(currentPage + 1)
    }} />)
    items.push(<Pagination.Last key = {totalPage + 2}  onClick={() => changePage(totalPage)} />)
    return (
        <Pagination className="justify-content-center mb-4" bsPrefix='pagination'>
            {items}
        </Pagination>
    )
}
export default Paginate;
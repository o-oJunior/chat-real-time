import React from "react"

type Props = {
  currentPage: number
  totalPages: number
  handlePageChange: (currentPage: number) => void
}

const Pagination = ({ currentPage, totalPages, handlePageChange }: Props) => {
  let visiblePages = []

  const maxVisible = Math.min(5, totalPages)

  if (totalPages <= maxVisible) {
    visiblePages = Array.from({ length: totalPages }, (_, i) => i + 1)
  } else {
    let startPage = Math.max(1, currentPage - Math.floor((maxVisible - 1) / 2))
    let endPage = startPage + maxVisible - 1

    if (endPage > totalPages) {
      endPage = totalPages
      startPage = endPage - maxVisible + 1
    }
    const pages = Array.from({ length: endPage - startPage + 1 }, (_, i) => startPage + i)
    visiblePages = [...pages]
  }
  return (
    <>
      {currentPage && (
        <div className="flex justify-center space-x-2">
          <button
            onClick={() => handlePageChange(Math.floor(currentPage / 2))}
            disabled={currentPage === 1}
            className={`px-3 py-1 text-sm bg-gray-200 rounded ${
              currentPage === 1 ? "opacity-50 cursor-not-allowed" : "hover:bg-gray-300"
            }`}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
              stroke="currentColor"
              className="w-4 h-4 text-black-600 hover:text-black-800 fill-current"
            >
              <path
                strokeWidth="2"
                fill="currentColor"
                d="M41.4 233.4c-12.5 12.5-12.5 32.8 0 45.3l160 160c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L109.3 256 246.6 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0l-160 160zm352-160l-160 160c-12.5 12.5-12.5 32.8 0 45.3l160 160c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L301.3 256 438.6 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0z"
              />
            </svg>
          </button>
          <button
            onClick={() => handlePageChange(currentPage - 1)}
            disabled={currentPage === 1}
            className={`px-3 py-1 text-sm bg-gray-200 rounded ${
              currentPage === 1 ? "opacity-50 cursor-not-allowed" : "hover:bg-gray-300"
            }`}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 320 512"
              stroke="currentColor"
              className="w-4 h-4 text-black-600 hover:text-black-800 fill-current"
            >
              <path
                strokeWidth="2"
                fill="currentColor"
                d="M41.4 233.4c-12.5 12.5-12.5 32.8 0 45.3l160 160c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L109.3 256 246.6 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0l-160 160z"
              />
            </svg>
          </button>

          {visiblePages.map((pageNumber, index) => {
            return (
              <React.Fragment key={pageNumber}>
                {index > 0 && visiblePages[index - 1] !== pageNumber - 1 && (
                  <span className="px-2 bg-gray-200 hover:bg-gray-300">...</span>
                )}
                <button
                  onClick={() => handlePageChange(pageNumber)}
                  className={`px-3 py-1 text-sm rounded ${
                    currentPage === pageNumber ? "bg-blue-500 text-white" : "bg-gray-200 hover:bg-gray-300"
                  }`}
                >
                  {pageNumber}
                </button>
              </React.Fragment>
            )
          })}

          <button
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
            className={`px-3 py-1 text-sm bg-gray-200 rounded ${
              currentPage === totalPages ? "opacity-50 cursor-not-allowed" : "hover:bg-gray-300"
            }`}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 320 512"
              stroke="currentColor"
              className="w-4 h-4 text-black-600 hover:text-black-800 fill-current"
            >
              <path
                strokeWidth="2"
                fill="currentColor"
                d="M278.6 233.4c12.5 12.5 12.5 32.8 0 45.3l-160 160c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L210.7 256 73.4 118.6c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l160 160z"
              />
            </svg>
          </button>
          <button
            onClick={() =>
              handlePageChange(Math.ceil(currentPage * 2 > totalPages ? totalPages : currentPage * 2))
            }
            disabled={currentPage === totalPages}
            className={`px-3 py-1 text-sm bg-gray-200 rounded ${
              currentPage === totalPages ? "opacity-50 cursor-not-allowed" : "hover:bg-gray-300"
            }`}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
              stroke="currentColor"
              className="w-4 h-4 text-black-600 hover:text-black-800 fill-current"
            >
              <path
                strokeWidth="2"
                fill="currentColor"
                d="M470.6 278.6c12.5-12.5 12.5-32.8 0-45.3l-160-160c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L402.7 256 265.4 393.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0l160-160zm-352 160l160-160c12.5-12.5 12.5-32.8 0-45.3l-160-160c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L210.7 256 73.4 393.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0z"
              />
            </svg>
          </button>
        </div>
      )}
    </>
  )
}

export default Pagination

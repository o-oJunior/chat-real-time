import { ChangeEvent, FormEvent, useState } from "react"

type Props = {
  handleChangeInput: (event: ChangeEvent<HTMLInputElement>) => void
  query: string
  textPlaceholder: string
  nameInput?: string
}
export default function Search({ query, textPlaceholder, handleChangeInput, nameInput }: Props) {
  return (
    <form className="flex items-center max-w-lg w-full">
      <input
        type="text"
        value={query}
        name={nameInput}
        onChange={handleChangeInput}
        maxLength={20}
        placeholder={`Pesquisar ${textPlaceholder}...`}
        className="w-full p-2 border border-gray-300 rounded-l-md focus:outline-none focus:ring-1 focus:ring-gray-400"
      />
      <button
        type="submit"
        className="p-3 bg-gray-500 text-white rounded-r-md hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-400"
      >
        <svg
          className="w-5 h-5"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M21 21l-4.35-4.35M18.5 10.5a7.5 7.5 0 1 1-15 0 7.5 7.5 0 0 1 15 0z"
          />
        </svg>
      </button>
    </form>
  )
}

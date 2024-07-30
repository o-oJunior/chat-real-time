import { useState } from "react"

interface IProps extends React.InputHTMLAttributes<HTMLInputElement> {
  type: string
}

const Input = ({ type = "text", ...props }: IProps) => {
  const [isFocused, setIsFocused] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const isTypePassword = type === "password"

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword)
  }

  return (
    <div className="relative mt-1 w-full">
      <input
        type={isTypePassword && showPassword ? "text" : type}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        {...props}
        className="block w-full px-3 pt-6 pb-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:border-gray-400 sm:text-sm placeholder-transparent"
      />
      {isTypePassword && (
        <>
          <svg
            onClick={togglePasswordVisibility}
            xmlns="http://www.w3.org/2000/svg"
            className={`absolute right-3 top-4 h-5 w-5 ${
              showPassword ? "hidden" : "block"
            } text-gray-500 cursor-pointer`}
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
            />
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M2.458 12C3.732 7.943 7.523 5 12 5c4.477 0 8.268 2.943 9.542 7-1.274 4.057-5.065 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
            />
          </svg>
          <svg
            onClick={togglePasswordVisibility}
            xmlns="http://www.w3.org/2000/svg"
            className={`absolute right-3 top-4 h-5 w-5 ${
              showPassword ? "block" : "hidden"
            } text-gray-500 cursor-pointer`}
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M13.875 18.825A10.05 10.05 0 0112 19c-4.477 0-8.268-2.943-9.542-7a10.05 10.05 0 012.42-3.988M4.1 4.1l15.8 15.8"
            />
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
            />
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M2.458 12C3.732 7.943 7.523 5 12 5c4.477 0 8.268 2.943 9.542 7-1.274 4.057-5.065 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
            />
          </svg>
        </>
      )}
      <label
        htmlFor={props.id || props.name}
        className={`absolute left-3 top-4 text-gray-500 pointer-events-none transform transition-all duration-200
          ${props.value || isFocused ? "top-0 -translate-y-3 text-sm" : "top-2 text-base"}
        `}
      >
        {props.placeholder}
      </label>
    </div>
  )
}

export default Input

import { createTheme } from "@mui/material";
import { Roboto } from "@next/font/google";

export const roboto = Roboto({
  weight: ["300", "400", "500", "700"],
  subsets: ["latin"],
  display: "swap",
  fallback: ["Helvetica", "Arial", "sans-serif"],
});

const theme = createTheme({
  palette: {
    primary: {
      main: "#B8C858",
    },
    secondary: {
      main: "#0E4987",
    },
    divider: '#1B73A7',
    background: {
      default: "#09345e",
      paper: "#0E4987",
    },
    mode: "dark",
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        "html, body, body>div": {
          padding: 0,
          margin: 0,
          width: "100%",
          height: "100%",
        },
        body: {
          background: "url(/img/background.png) no-repeat center center fixed",
          WebkitBackgroundSize: "cover",
          OBackgroundSize: "cover",
          BackgroundSize: "cover",
          MozBackgroundSize: "cover",
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: ({ theme, ownerState }) => ({
          ...(ownerState.variant === "outlined" && {
            border: `3px solid ${theme.palette.divider}`,
          }),
        }),
      },
    },
  },
});

export default theme;

import { useState, useEffect } from "react";
import { useTheme } from "next-themes";
import {Switch} from "@nextui-org/react";


export const ThemeSwitcher = () => {
  const [mounted, setMounted] = useState(false);
  const { theme, setTheme } = useTheme();
  const [isSelected, setIsSelected] = useState(theme === "dark");

  useEffect(() => {
    setMounted(true);
  }, []);


  if (!mounted) {
    return null;
  }

  const handleValueChange = (selected: boolean) => {
    if(selected){
      setTheme("dark");
    } else {
      setTheme("light");
    }
    setIsSelected(selected);
  }


  return (
    <>
      <Switch isSelected={isSelected} onValueChange={handleValueChange}>
          {isSelected ? "Dark Mode" : "Light Mode"}
      </Switch>  
    </>
    );

};
